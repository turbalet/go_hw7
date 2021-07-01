package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

type Server struct {
	listener net.Listener
	quit     chan interface{}
	wg       sync.WaitGroup
	sm       *Semaphore
}

func NewServer(addr string, size int) *Server {
	s := &Server{
		quit: make(chan interface{}),
		sm:   NewSemaphore(size),
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s.listener = l
	s.wg.Add(1)
	go s.serve()
	return s
}

func (s *Server) Stop() {
	close(s.quit)
	// stops accepting new connection -> will cause error in serve() when Accept() -> return
	s.listener.Close()
	fmt.Println("Stop accepting new connections")
	// waits until open connections returns
	s.wg.Wait()
	fmt.Println("All connections are closed")
}

func (s *Server) serve() {
	defer s.wg.Done()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				return
			default:
				fmt.Printf("accept error: %v\n", err)
			}
		} else {
			s.wg.Add(1)
			s.sm.Acquire(1)
			go func() {
				defer s.wg.Done()
				defer s.sm.Release(1)
				s.handleConnection(conn)
			}()
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("error reading:", err.Error())
	}

	strVal := string(buf[:n])
	num, err := strconv.Atoi(strVal)

	if err != nil {
		fmt.Printf("Failed to convert %s to integer\n", strVal)
	}

	conn.Write([]byte(strconv.Itoa(num * num)))
}
