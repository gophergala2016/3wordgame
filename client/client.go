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
		os.Exit(-1)
	}

	clear("Connected.")

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading from Stdin.")
		}

		fmt.Fprintf(conn, input)
	}

	fmt.Println("Exiting.")
}

func clear(msg string) {
	print("\033[H\033[2J")
	fmt.Println(msg)
	fmt.Print("Input > ")
}
