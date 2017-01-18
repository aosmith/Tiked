package main


var MagicNumber int64 = 0;  

func BypassAV() {
  LongLoop()
  Jump()
  Wait()
  //CheckDebugger()
}


func Jump() {
  MagicNumber++
  hop1()
}

func LongLoop()  {
  for i := 0; i < 1000000; i++ {}
}

func CheckDebugger() {
  Flag,_,_ := IsDebuggerPresent.Call()
  if Flag != 0 {
    Wait()
    CheckDebugger()
  }
}

func hop1() {
  MagicNumber++
  hop2()
}
func hop2() {
  MagicNumber++
  hop3()
}
func hop3() {
  MagicNumber++
  hop4()
}
func hop4() {
  MagicNumber++
  hop5()
}
func hop5() {
  MagicNumber++
  hop6()
}
func hop6() {
  MagicNumber++
  hop7()
}
func hop7() {
  MagicNumber++
  hop8()
}
func hop8() {
  MagicNumber++
  hop9()
}
func hop9() {
  MagicNumber++
  hop10()
}
func hop10() {
  MagicNumber++
}
