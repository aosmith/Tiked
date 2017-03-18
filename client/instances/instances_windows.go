package instances

import "os"
import "github.com/rodolfoag/gow32"

func CheckMultiInstances() {
	_, err := gow32.CreateMutex("Windows_Security")
	if err != nil {
		// TODO start in wait mode
		//Run("msg * Running")
		os.Exit(0)

	}
}
