package main

import "os/exec"

func Run(cmd string) string {
	e := exec.Command("/bin/sh", "-c", cmd)
	res, _ := e.Output()
	return string(res)
}
