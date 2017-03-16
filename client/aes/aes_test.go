package aes

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	key_text = make([]byte, 32) // Rand 32 bytes

	input1 := make([]byte, 1028)
	input2 := make([]byte, 1028)
	input3 := make([]byte, 1028)

	_, err := rand.Read(key_text)
	_, err = rand.Read(input1)
	_, err = rand.Read(input2)
	_, err = rand.Read(input3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

}
