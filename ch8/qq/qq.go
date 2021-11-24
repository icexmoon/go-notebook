package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
		return
	}
	go createChatRoom()
	fmt.Println("chat room is created in localhost:8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// fmt.Println("get a request from " + conn.RemoteAddr().String())
		go handleConn(conn)
	}
}

type clientChan chan string

var enterChan = make(chan clientChan)
var exitChan = make(chan clientChan)
var msgChan = make(chan string)

func createChatRoom() {
	clients := make(map[clientChan]bool)
	for {
		select {
		case newClientChan := <-enterChan:
			clients[newClientChan] = true
		case leavedClientChan := <-exitChan:
			delete(clients, leavedClientChan)
			close(leavedClientChan)
		case msg := <-msgChan:
			for cChan := range clients {
				cChan <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	// log.Println("handle " + conn.RemoteAddr().String())
	chatChan := make(clientChan)
	name := conn.RemoteAddr().String()
	msg := name + " is joined this chat rom."
	msgChan <- msg
	log.Println(msg)
	enterChan <- chatChan
	go writeClientChanToClient(conn, chatChan)
	readClientToMsgChan(conn, name)
	exitChan <- chatChan
	msg = name + " is leaved."
	msgChan <- msg
	log.Println(msg)
}

func readClientToMsgChan(conn net.Conn, name string) {
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		msgChan <- name + " said: " + sc.Text()
	}
}

func writeClientChanToClient(conn net.Conn, chatChan clientChan) {
	for msg := range chatChan {
		fmt.Fprintln(conn, msg)
	}
}
