all: client server frontend

client: client/*.go *.go
	go build -o bin/client client/client.go

frontend: frontend.go *.go
	go build -o bin/frontend client/frontend.go

server: server/*.go *.go
	go build -o bin/server server/server.go

clean:
	rm -rf bin/client bin/frontend bin/server
