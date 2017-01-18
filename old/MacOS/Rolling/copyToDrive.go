package main

//import "github.com/cleversoap/go-cp"
import "os"
import "strings"

func GetExeName() string {
	filename := strings.Split(os.Args[0], "\\")
	return filename[len(filename)-1]
}

func GetParentFolder() string {
	filename := strings.Split(os.Args[0], "\\")
	return filename[len(filename)-2]
}