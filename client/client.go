package client

import (
	"fmt"
	"net"
	"os"
	"time"
)

func RunClient() {
	strEcho := "5"
	servAddr := "localhost:3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		fmt.Println("Failed to write:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Written = ", strEcho)
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		fmt.Println("Failed to write", err.Error())
		os.Exit(1)
	}

	fmt.Println("Received = ", string(reply))

}

func RunClientWithDelay() {
	strEcho := "5"
	servAddr := "localhost:3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	time.Sleep(time.Millisecond * 1000)

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		fmt.Println("Failed to write:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Written = ", strEcho)
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		fmt.Println("Failed to write", err.Error())
		os.Exit(1)
	}

	fmt.Println("Received = ", string(reply))

}
