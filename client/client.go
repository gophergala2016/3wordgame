package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

// var inputChannel = make(chan string)

func main() {
	var server string
	var port int

	flag.StringVar(&server, "server", "127.0.0.1", "Server host")
	flag.IntVar(&port, "port", 6666, "Server port")
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		fmt.Println("Error dialing in.")
		exit()
	}

	clear("Connected.")

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from Stdin.")
			exit()
		}

		fmt.Fprintf(conn, input)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from conn.")
			exit()
		}

		fmt.Println(fmt.Sprintf("Read: %s", message))
	}

	fmt.Println("Exiting.")
}

func clear(msg string) {
	print("\033[H\033[2J")
	fmt.Println(msg)
	fmt.Print("Input > ")
}

func exit() {
	os.Exit(-1)
}
