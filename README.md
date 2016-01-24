# 3wordgame

Collaboratively write a story, 3 words at a time! The only trick... you can't follow your own 3 words.

## Quick Start

```sh
$ git clone https://github.com/gophergala2016/3wordgame.git
$ go run client.go
```

### Hosting
To run your own server, simply:

```sh
$ go run server.go
```

By default the server listens on port 8080, though this can be changed by specifing the `--port` option e.g.:

```sh
$ go run server.go --port 8080
```

Clients by default try to connect to `3wordgame.com:8080` this can also be changed by specifying `--server` and `--port` args e.g.:

```sh
$ go run client.go --server 127.0.01 --port 9090
```

## Contributing

To run tests for the validation

```sh
$ go test ./validation
```
