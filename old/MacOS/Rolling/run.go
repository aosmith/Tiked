package main

import "os/exec"

//Start command in minimized window and returns output

func Run(cmd string) string {
	e := exec.Command(cmd)
	res, _ := e.Output()
	e.Run()
	
	return string(res)
}
