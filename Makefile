all: client server

client: 
	go build -o bin/client client.go

server:
	go build -o bin/server server.go

clean:
	rm -rf bin/client bin/frontend bin/server
