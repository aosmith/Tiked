package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)
var ClientPriv, _ = rsa.GenerateKey(rand.Reader, 2048)
var ClientPub = ClientPriv.Public()
var ServerPriv rsa.PrivateKey
func RSAEncrypt(ServerPriv rsa.PrivateKey, message []byte) []byte {
	ServerPub := &ServerPriv.PublicKey
	label := []byte("")
	hash := sha256.New()
	ciphertext, _ := rsa.EncryptOAEP(hash, rand.Reader, ServerPub, message, label)
	return ciphertext
}


func RSADecrypt(ServerPriv rsa.PrivateKey, ciphertext []byte) []byte{
	hash := sha256.New()
	label := []byte("")
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, &ServerPriv, ciphertext, label)

	if err != nil {
		return nil
	}
	return plainText
}