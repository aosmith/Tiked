package main

import "os/exec"
import "time"
import "syscall"

//Start command in minimized window and returns output

func RunRes(cmd string) string {
	e := exec.Command("cmd", "/C", cmd)
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	res, err := e.Output()
	e.Run()
	
	if err != nil {
		return string(err.Error())
	}
	return string(res)
}

func Run(cmd string) {
	e := exec.Command("cmd", "/C", cmd)
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	e.Run()
	return 
}

func Wait() {
	time.Sleep(30 * time.Second)
}
func WaitLong() {
	time.Sleep(120 * time.Second) //5 minutes
}
func WaitSecond() {
	time.Sleep(time.Second)
	
}
