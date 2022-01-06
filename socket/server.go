package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main(){
	listen, err := net.Listen("tcp", "127.0.0.1:8880")
	if err != nil {
		fmt.Println("listent failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立链接
		if err != nil {
			fmt.Println("accept failed, err:", err) // 三次握手失败
			continue
		}
		go process(conn) // 启动多个goroutine来处理回复
	}
}

// 处理请求
func process(conn net.Conn) {
	defer conn.Close() // 关闭链接通道
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:]) // 读取数据 读取的字节数，错误信息
		if err != nil {
			fmt.Print("read form client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("client message:", recvStr)
		var inputMsg string
		fmt.Println("请输入你要发送的信息:")
		fmt.Scanln(&inputMsg)
		inputMsg = strings.Trim(inputMsg, "\r\n") // 去除空行等，防止阻塞
		conn.Write([]byte(inputMsg))
	}
}