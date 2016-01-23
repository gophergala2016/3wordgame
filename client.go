package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

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

	fmt.Println("Connected.")

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Printf("Error reading from Stdin.")
		}

		fmt.Fprintf(conn, input)
	}

	fmt.Println("Exiting!")
}
