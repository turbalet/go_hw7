package main

import (
	"fmt"
	"puppy/client"
	"puppy/server"
	"time"
)

func main() {
	s := server.NewServer(":3333", 3)
	// to test semaphore
	//go client.RunClientWithDelay()
	//go client.RunClientWithDelay()
	// runs with delay
	go client.RunClientWithDelay()
	// runs сразу
	go client.RunClient()
	go client.RunClient()
	time.Sleep(time.Millisecond * 100)
	s.Stop()
	fmt.Println("Server stopped. Closed for new connections")
	// Expecting fail.
	go client.RunClient()
	time.Sleep(time.Second * 5)
}
