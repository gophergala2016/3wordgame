package main

import (
	"bufio"
	"net"
	"fmt"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":6666")
	if err != nil {
		fmt.Println("Error dialing in.");
	}

	fmt.Println("Connected.");

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Printf("Error reading from Stdin.")
		}

		fmt.Fprintf(conn, input)
	}

	fmt.Println("Exiting!")
}
