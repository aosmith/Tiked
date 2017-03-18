package utils

import "time"

// Wait waits 30 seconds
func Wait(s time.Duration) {
	time.Sleep(s * time.Second)
}
