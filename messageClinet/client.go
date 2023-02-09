package main

import (
	"Simple-Douyin-Backend/messageServer"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func login(conn net.Conn, userId, toUserId int64) {
	req := messageServer.SendEvent{UserId: userId, ToUserId: toUserId, MsgContent: ""}
	reqData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Marshal error")
	}
	fmt.Printf("%d", len(reqData))
	if _, err := conn.Write(reqData); err != nil {
		fmt.Println("login error")
	}
	fmt.Println("login over")
}

func sender(conn net.Conn, userId, toUserId int64, content string) {
	req := messageServer.SendEvent{UserId: userId, ToUserId: toUserId, MsgContent: content}
	reqData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("send message format error")
	}
	if _, err := conn.Write(reqData); err != nil {
		fmt.Println("send error")
	}
	fmt.Println("send over")

	//接收服务端反馈
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
		return
	}
	fmt.Println(conn.RemoteAddr().String(), "receive server back msg: ", string(buffer[:n]))
}

func main() {
	server := "127.0.0.1:9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	login(conn, 1, 3)

	time.Sleep(time.Second)

	login(conn, 3, 1)

	time.Sleep(time.Second)
	sender(conn, 1, 3, "hello")

}
