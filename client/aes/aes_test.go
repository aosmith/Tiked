package aes

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"../base64"
	"../utils"
	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	tempDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	keyText := make([]byte, 32) // Rand 32 bytes
	_, _ = rand.Read(keyText)
	myIv := make([]byte, 16) // Rand 32 bytes
	_, _ = rand.Read(myIv)

	InitializeBlock(keyText, myIv, "irrelevant")
	fmt.Println("iv: ", base64.Base64EncodeRaw(iv[:]))
	fmt.Println("keyText: ", base64.Base64EncodeRaw(keyText))

	input1 := []byte("ooooooooooooooooooooo")

	// Write some files
	err := ioutil.WriteFile(tempDir+"/dat1", input1, 0777)
	utils.CheckT(err, t)

	i1, _ := ioutil.ReadFile(tempDir + "/dat1")
	err = StreamEncrypter(tempDir + "/dat1")
	if err != nil {
		panic(err)
	}

	// Decrypt
	err = StreamDecrypter(tempDir + string(os.PathSeparator) + /*base64.Base64Encode("dat1")*/ "dat1" /* + Ext*/)
	if err != nil {
		panic(err)
	}

	f1, _ := ioutil.ReadFile(tempDir + "/dat1")
	assert.Equal(t, string(i1), strings.Trim(string(f1), " "), "should be equal")
}
