package utils

import (
	"fmt"
	"testing"
)

// CheckT recives error, *testing.T and logs any errors
func CheckT(err error, t *testing.T) {
	if err != nil {
		fmt.Println("error:", err)
		t.Error()
		return
	}
}

// Check recives and error and panics
func Check(err error) {
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}
