package utils

import (
	"fmt"
	"net"
	"strings"
)

func GetOutBoundIP() (ip string, err error) {
	// 使用udp发起网络连接, 这样不需要关注连接是否可通, 随便填一个即可
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
