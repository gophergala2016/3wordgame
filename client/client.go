package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gophergala2016/3wordgame/validation"
	"net"
	"os"
	"strings"
	"time"
)

var story string

func Read(conn net.Conn) {
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

func Write(conn net.Conn) {
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

func Listen(conn net.Conn) {
	go Read(conn)
	Write(conn)
}

func main() {
	var server string
	var port int

	flag.StringVar(&server, "server", "127.0.0.1", "Server host")
	flag.IntVar(&port, "port", 6666, "Server port")
	flag.Parse()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", server, port), time.Second*3)
	if err != nil {
		fmt.Println("Error dialing in.")
		exit()
	}

	clear("Connected.")

	Listen(conn)

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
