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
	incoming   chan Message
	outgoing   chan string
	reader     *bufio.Reader
	writer     *bufio.Writer
	address    string
	connection net.Conn
}

type Message struct {
	text    string
	address string
}

// Read line by line into the client.incoming
func (client *Client) Read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err != nil {
			client.connection.Close()
			fmt.Println(fmt.Sprintf("Read error %s", err))
			break
		}

		client.incoming <- Message{
			text:    line,
			address: client.address,
		}

		fmt.Println(fmt.Sprintf("Client.Read %s", line))
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

	address := connection.RemoteAddr().String()

	client := &Client{
		incoming:   make(chan Message),
		outgoing:   make(chan string),
		reader:     reader,
		writer:     writer,
		address:    address,
		connection: connection,
	}

	client.Listen()

	fmt.Println("NewClient")

	return client
}

// ChatRoom struct
type ChatRoom struct {
	clients               []*Client
	joins                 chan net.Conn
	incoming              chan Message
	outgoing              chan string
	story                 string
	last_msg_user_address string
}

// Broadcast data to all connected chatRoom.clients
func (chatRoom *ChatRoom) Broadcast(data string) {
	for _, client := range chatRoom.clients {
		client.outgoing <- fmt.Sprintf("%s\n", data)
	}
}

// Join attaches a new client to the chatRoom clients
func (chatRoom *ChatRoom) Join(connection net.Conn) {
	client := NewClient(connection)
	chatRoom.clients = append(chatRoom.clients, client)
	client.outgoing <- fmt.Sprintf("%s\n", chatRoom.story)
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
				msg, err := validation.ValidateMsg(data.text)
				if err == nil && chatRoom.last_msg_user_address != data.address {
					fmt.Println(fmt.Sprintf("chatRoom.Broadcast %s", msg))

					chatRoom.Broadcast(msg)
					if len(chatRoom.story) == 0 {
						chatRoom.story = msg
					} else {
						chatRoom.story = fmt.Sprintf("%s %s", chatRoom.story, msg)
					}
					chatRoom.last_msg_user_address = data.address
				}
			case conn := <-chatRoom.joins:
				fmt.Println("chatRoom.join")
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
		incoming: make(chan Message),
		outgoing: make(chan string),
	}

	chatRoom.Listen()

	fmt.Println("NewChatRoom")

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
