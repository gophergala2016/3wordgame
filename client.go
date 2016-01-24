package main

import (
	"./validation"
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var story string

func read(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from conn.")
			exit()
		}

		story = fmt.Sprintf("%s %s", story, validation.StripNewLine(message))
		story = strings.Trim(story, " \n")

		showStory()
	}
}

func write(conn net.Conn) {
	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from Stdin.")
			exit()
		}
		fmt.Fprintf(conn, input)
		showStory()
	}
}

func listen(conn net.Conn) {
	go read(conn)
	write(conn)
}

func main() {
	var server string
	var port int

	flag.StringVar(&server, "server", "3wordgame.com", "Server host")
	flag.IntVar(&port, "port", 8080, "Server port")
	flag.Parse()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", server, port), time.Second*3)
	if err != nil {
		fmt.Println("Error dialing in.")
		exit()
	}

	clear("Connected.")

	listen(conn)

	fmt.Println("Exiting.")
}

func clear(msg string) {
	print("\033[H\033[2J")
	fmt.Println(msg)
	fmt.Print("Input > ")
}

func showStory() {
	if len(story) == 0 {
		clear(fmt.Sprintf("No story yet..."))
	} else {
		clear(fmt.Sprintf("Story so far: %s", story))
	}
}

func exit() {
	os.Exit(-1)
}
