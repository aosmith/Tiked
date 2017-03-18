package aes

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"../base64"
	"../utils"
	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	tempDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	key_text := make([]byte, 32) // Rand 32 bytes
	_, _ = rand.Read(key_text)

	InitializeBlock(key_text, "ghkkfiki")

	input1 := make([]byte, 1028)
	input2 := make([]byte, 1028)
	input3 := make([]byte, 1028)

	_, err := rand.Read(key_text)
	utils.CheckT(err, t)
	_, err = rand.Read(input1)
	utils.CheckT(err, t)
	_, err = rand.Read(input2)
	utils.CheckT(err, t)
	_, err = rand.Read(input3)
	utils.CheckT(err, t)

	// Write some files
	err = ioutil.WriteFile(tempDir+"/dat1", input1, 0777)
	utils.CheckT(err, t)
	err = ioutil.WriteFile(tempDir+"/dat2", input2, 0777)
	utils.CheckT(err, t)
	err = ioutil.WriteFile(tempDir+"/dat3", input3, 0777)
	utils.CheckT(err, t)

	i1, _ := ioutil.ReadFile(tempDir + "/dat1")
	i2, _ := ioutil.ReadFile(tempDir + "/dat2")
	i3, _ := ioutil.ReadFile(tempDir + "/dat3")

	err = StreamEncrypter(tempDir + "/dat1")
	utils.CheckT(err, t)
	err = StreamEncrypter(tempDir + "/dat2")
	utils.CheckT(err, t)
	err = StreamEncrypter(tempDir + "/dat3")
	utils.CheckT(err, t)

	// Print input and out
	bi1 := base64.Base64EncodeRaw(i1[:16])
	bi2 := base64.Base64EncodeRaw(i2[:16])
	bi3 := base64.Base64EncodeRaw(i3[:16])
	fmt.Println("In:")
	fmt.Println(bi1)
	fmt.Println(bi2)
	fmt.Println(bi3)
	fmt.Println()

	fmt.Println("Out:")

	c1, err := ioutil.ReadFile(tempDir + string(os.PathSeparator) + base64.Base64Encode("dat1") + Ext)
	c2, _ := ioutil.ReadFile(tempDir + string(os.PathSeparator) + base64.Base64Encode("dat3") + Ext)
	c3, _ := ioutil.ReadFile(tempDir + string(os.PathSeparator) + base64.Base64Encode("dat2") + Ext)
	bo1 := base64.Base64EncodeRaw(c1[:16])
	bo2 := base64.Base64EncodeRaw(c2[:16])
	bo3 := base64.Base64EncodeRaw(c3[:16])
	fmt.Println(bo1)
	fmt.Println(bo2)
	fmt.Println(bo3)

	assert.NotEqual(t, bi1, bo1, "they should not be equal")
	assert.NotEqual(t, bi2, bo2, "they should not be equal")
	assert.NotEqual(t, bi3, bo3, "they should not be equal")

	// Decrypt
	StreamDecrypter(tempDir + string(os.PathSeparator) + base64.Base64Encode("dat1") + Ext)
	StreamDecrypter(tempDir + string(os.PathSeparator) + base64.Base64Encode("dat2") + Ext)
	StreamDecrypter(tempDir + string(os.PathSeparator) + base64.Base64Encode("dat3") + Ext)

	//f1, _ := ioutil.ReadFile(tempDir + string(os.PathSeparator) + "dat1")
	//f2, _ := ioutil.ReadFile(tempDir + string(os.PathSeparator) + "dat2")
	//f3, _ := ioutil.ReadFile(tempDir + string(os.PathSeparator) + "dat3")

	fmt.Println(string(i1))
	//assert.Equal(t, i1, f1, "should be equal")
	//assert.Equal(t, input2, f2, "should be equal")
	//assert.Equal(t, input3, f3, "should be equal")

}
