// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	closeChan := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("server print is closed")
		closeChan <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-closeChan
	log.Println("user input is closed")
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
