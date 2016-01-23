package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gophergala2016/3wordgame"
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
		threewordgame.SetStatus(fmt.Sprintf("Error dialing in."))
		threewordgame.Exit()
		os.Exit(-1)
	}

	threewordgame.SetStatus(fmt.Sprintf("Connected."))

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			threewordgame.SetStatus(fmt.Sprintf("Error reading from Stdin."))
		}

		fmt.Fprintf(conn, input)
	}

	threewordgame.SetStatus(fmt.Sprintf("Exiting!"))
}
