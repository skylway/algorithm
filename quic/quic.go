package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"

	"github.com/lucas-clemente/quic-go"
)

const addr = "localhost:4242"

const clientMessage = "Hello"
const serverMessage = "Hi"

func main() {
	go func() {
		err := server()
		if err != nil {
			panic(err)
		}
	}()
	err := client()
	if err != nil {
		panic(err)
	}

	// 等待main和go程 执行完,防止server执行完自动结束
	time.Sleep(time.Second * 50)
}

// 客户端
func client() error {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}

	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		return err
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Client: Sending '%s'\n", clientMessage)
	_, err = stream.Write([]byte(clientMessage))
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	_, err = stream.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("Client: Got '%s'\n", buf)

	return nil
}

// 服务端
func server() error {
	tls :=  generateTLSConfig()
	fmt.Println("tls:", tls)
	listener, err := quic.ListenAddr(addr, tls, nil)
	if err != nil {
		return err
	}
	sess, err := listener.Accept(context.Background())
	if err != nil {
		return err
	}
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	_, err = stream.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("Server: Got '%s'\n", buf)

	fmt.Printf("Server: Sending '%s'\n", serverMessage)
	_, err = stream.Write([]byte(serverMessage))
	return err
}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	fmt.Println("key:", key)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	fmt.Println("template:", template)

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("certDER:", certDER)

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	fmt.Println("tlsCert:", tlsCert)

	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}