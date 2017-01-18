package main

import (
	"fmt"
	"net/http"
)


//Start .jar

func main() {  //Never ends
	fmt.Println("Listening for token...")
	http.HandleFunc("/", handleOAuth2Callback)
	http.ListenAndServe(":1337", nil)
}

func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	//Save code to database
	fmt.Println(code)
}

//Go to this url and accept
//https://accounts.google.com/o/oauth2/auth?access_type&approval_prompt&scope=https%3A%2F%2Fwww.google.com%2Fm8%2Ffeeds&response_type=code&redirect_uri=http://127.0.0.1:1337&client_id=36575666534-pjri8bbmod5a3onnhd1j6tqimhjt0sfc.apps.googleusercontent.com&from_login=1&as=-7947c60104d4829d&authuser=0
