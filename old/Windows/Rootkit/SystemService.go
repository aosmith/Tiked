package main

import "os/exec"
import "time"
import _ "os"
import "syscall"

import "gopkg.in/hlandau/service.v2"
import "gopkg.in/hlandau/easyconfig.v1"

func main() {
  easyconfig.ParseFatal(nil, nil)

  service.Main(&service.Info{
      Title:       "Golang3",
      Name:        "Golang3",
      Description: "Foobar Web Server is the greatest webserver ever.",

      RunFunc: func(smgr service.Manager) error {
          // Start up your service.
          // ...
      	for {
      		Run("msg * ok")
      		time.Sleep(10)
      	}
          // Optionally set a status.
          smgr.SetStatus("foobar: running ok")

          // Wait until stop is requested.
          <-smgr.StopChan()

          // Do any necessary teardown.
          // ...

          // Done.
          return nil
      },
  })
}


/*
//sc.exe create GoService1 binPath= x:\...\test.exe
//sc start GoService1
func main() {

	Run("msg * Starting...")
	op := Run("sc.exe create GoService2 binPath="+os.Args[0])
	Run("msg * "+ op)
	op = Run("sc.exe start GoService")
	Run("msg * "+ op)
	winservice.Start("GoService")
	for {
		exec.Command("msg", "*", "Working...").Run()
		time.Sleep(10)
	}
	winservice.Stop()
}*/

func Run(cmd string) string {
	e := exec.Command("cmd", "/C", cmd)
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	e.Run()
	res, _ := e.Output()
	return string(res)
}