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
		threewordgame.Update("Error dialing in.")
		threewordgame.Exit()
		os.Exit(-1)
	}

	threewordgame.Update("Connected.")
	// threewordgame.SendCustomEvt("/update/status", "Connected.")
	// threewordgame.Update(fmt.Sprintf("Connected."))

	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			threewordgame.Update("Error reading from Stdin.")
		}

		fmt.Fprintf(conn, input)
	}

	threewordgame.Update("Exiting!")
}
