package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println(msg)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	//fmt.Println("omatch")
	for {
		//fmt.Printf("Enter message -> ")
		msg, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error writing to server:", err)
			return
		}
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server

	conn, err := net.Dial("tcp", *addrPtr)
	//fmt.Println("omash")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	//fmt.Println("super")
	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	//TODO Start getting and sending user messages.
	write(conn)

	//fmt.Println("shambo")
}
