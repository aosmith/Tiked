package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "encoding/pem"
    "crypto/x509"
    "io/ioutil"
    "encoding/base64"
    "fmt"
)
/*
func main() {
    dat, _ := ioutil.ReadFile("pub")
    pub := LoadPubKey(dat)
    fmt.Println("g")
    dat, _ = ioutil.ReadFile("priv")
    priv := LoadPrivKey(dat)

    chiper, _ := RSAEncrypt([]byte("Hola"), &pub.PublicKey)
    fmt.Println(string(chiper))
    clear, _ := RSADecrypt(chiper,&priv)
    fmt.Println(string(clear))
}*/


func LoadPrivKey(dat []byte) rsa.PrivateKey {
    var PrivKey rsa.PrivateKey
    temp, err := x509.ParsePKCS1PrivateKey(dat)
    if err != nil {fmt.Println(err.Error());}
    PrivKey = *temp
    return PrivKey;
}

func LoadPubKey(dat []byte) rsa.PrivateKey {
    var PubKey rsa.PrivateKey
    temp, err := x509.ParsePKIXPublicKey(dat)
    if err != nil {fmt.Println("Pub error: ",err.Error());}
    temp2 := temp.(*rsa.PublicKey)
    PubKey.PublicKey = *temp2
    return PubKey;
}

func GenerateRSA() {
    PrivateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    PublicKey := &PrivateKey.PublicKey
    pub, _ := x509.MarshalPKIXPublicKey(PublicKey)
    ioutil.WriteFile("pub", []byte(pub), 0644)
    ioutil.WriteFile("priv", []byte(x509.MarshalPKCS1PrivateKey(PrivateKey)), 0644)
}

func RSAEncrypt(clearText []byte, pub *rsa.PublicKey) ([]byte, error) {
    message := clearText
    label := []byte("")
    hash := sha256.New()
    return rsa.EncryptOAEP(hash, rand.Reader, pub, message, label)
}

func RSADecrypt(cipherText []byte, priv *rsa.PrivateKey) ([]byte, error){
    label := []byte("")
    hash := sha256.New()
    return rsa.DecryptOAEP(hash, rand.Reader, priv, cipherText, label)
}
func RSAPrivPem(Priv rsa.PrivateKey) string {
    pemdata := pem.EncodeToMemory(
        &pem.Block{
            Type: "RSA PRIVATE KEY",
            Bytes: x509.MarshalPKCS1PrivateKey(&Priv),
        },
    )
    return string(pemdata)
}
func Base64Encode(str []byte) string {
    return base64.StdEncoding.EncodeToString(str)
}

func Base64Decode(str string) []byte {
    data, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return nil
    }
    return data
}
