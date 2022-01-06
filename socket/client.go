package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999") // 绑定服务端地址
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close() // 关闭双向链接
	for {
		var inputMsg string
		fmt.Println("请输入你要发送的信息:")
		fmt.Scanln(&inputMsg)
		inputMsg = strings.Trim(inputMsg, "\r\n") // 去除空行等，防止阻塞
		if strings.ToUpper(inputMsg) == "quit" {
			return
		}
		_, err = conn.Write([]byte(inputMsg)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		serverMsg, err := conn.Read(buf[:]) // 服务端返回的信息
		if err != nil {
			fmt.Println("recv failed err:", err)
			return
		}
		fmt.Println("server message:", string(buf[:serverMsg]))
	}
}

