package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const (
	addr = "127.0.0.1:1100"
)

func main() {
	DeadlineListener()
}

func Listener() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer func() {
				err := c.Close()
				if err != nil {
					log.Fatal(err)
				}
			}()

			buf := make([]byte, 1024)

			for {
				n, err := c.Read(buf)
				if err != nil {
					if err != io.EOF {
						log.Fatal(err)
					}
					return
				}
				fmt.Printf("received: %q", buf[:n])
			}
		}(conn)
	}
}

func DeadlineListener() {
	listener, err := net.Listen("tcp", addr)
	handleError(err)

	conn, err := listener.Accept()
	handleError(err)
	defer func() {
		err = conn.Close()
		handleError(err)
	}()
	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Println(err)
		return
	}
	buf := make([]byte, 1024)
	_, err = conn.Read(buf) // blocked until remote node sends data
	log.Println(string(buf))
	nErr, ok := err.(net.Error)
	if ok && nErr.Timeout() {
		log.Println(err)
	}
}

func HealthBeat() {

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
