package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gophergala2016/3wordgame/validation"
	"net"
)

// Client struct
type Client struct {
	incoming chan string
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

// Read line by line into the client.incoming
func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.incoming <- line
	}
}

// Write client outgoing data to the client writer
func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

// Listen for reads and writes on the client
func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

// NewClient returns new instance of client.
func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		incoming: make(chan string),
		outgoing: make(chan string),
		reader:   reader,
		writer:   writer,
	}

	client.Listen()

	return client
}

// ChatRoom struct
type ChatRoom struct {
	clients  []*Client
	joins    chan net.Conn
	incoming chan string
	outgoing chan string
}

// Broadcast data to all connected chatRoom.clients
func (chatRoom *ChatRoom) Broadcast(data string) {
	for _, client := range chatRoom.clients {
		client.outgoing <- data
	}
}

// Join attaches a new client to the chatRoom clients
func (chatRoom *ChatRoom) Join(connection net.Conn) {
	client := NewClient(connection)
	chatRoom.clients = append(chatRoom.clients, client)
	go func() {
		for {
			chatRoom.incoming <- <-client.incoming
		}
	}()
}

// Listen to all incoming messages for the chatRoom
func (chatRoom *ChatRoom) Listen() {
	go func() {
		for {
			select {
			case data := <-chatRoom.incoming:
				msg, err := validation.ValidateMsg(data)
				if err == nil {
					chatRoom.Broadcast(msg)
				}
			case conn := <-chatRoom.joins:
				chatRoom.Join(conn)
			}
		}
	}()
}

// NewChatRoom factories a ChatRoom instance
func NewChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		clients:  make([]*Client, 0),
		joins:    make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	chatRoom.Listen()

	return chatRoom
}

func main() {
	var server string
	var port int

	flag.StringVar(&server, "server", "", "Server host")
	flag.IntVar(&port, "port", 6666, "Server port")
	flag.Parse()

	listener, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", server, port))

	chatRoom := NewChatRoom()

	for {
		conn, _ := listener.Accept()
		chatRoom.joins <- conn
	}
}
