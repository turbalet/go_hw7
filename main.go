package main

import (
	"fmt"
	"puppy/client"
	"puppy/server"
	"time"
)

func main() {
	s := server.NewServer(":3333", 5)
	go client.RunClientWithDelay()
	go client.RunClientWithDelay()
	time.Sleep(time.Millisecond * 100)
	s.Stop()
	fmt.Println("Server stopped. Closed for new connections")
	go client.RunClient()
	time.Sleep(time.Second * 5)
}
