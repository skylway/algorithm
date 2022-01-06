package main

import (
	"log"
	"net"
)

func main() {
	go listenMaster(":4000") // 服务器主连接的端口号
	go listenMaster(":8000") // 服务器响应客户端打洞申请的端口号
	for {}
}

func listenMaster(port string) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("error listen:",port, err)
		return
	}
	defer l.Close()
	log.Println("listen ok", port)
	var i int
	for {
		// time.Sleep(time.Second * 10)
		c, err := l.Accept();
		if err != nil {
			log.Println("accept error:", port, err)
			break
		}
		i++
		log.Printf("%d: accept a new connection port %s", i, port)
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}