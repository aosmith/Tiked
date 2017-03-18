package utils

import (
	"fmt"
	"testing"
)

func CheckT(err error, t *testing.T) {
	if err != nil {
		fmt.Println("error:", err)
		t.Error()
		return
	}
}

func Check(err error) {
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}
