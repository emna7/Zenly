# Using GRPC with Golang for NTP server

Creating a server (and demo client) using Go and GRPC that responds to time requests and fetches
the time via NTP.

## Go Installation

After Downlowding Go:
1- Extract the archive downloaded into usr/local:
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz

2- Add /usr/local/go/bin to the PATH environment variable
export PATH=$PATH:/usr/local/go/bin

3- Verify that you've installed Go by opening a command prompt and typing the following command: 
$ go version

## Protocol Buffer Compiler Installation

$ apt install -y protobuf-compiler
$ protoc --version

## Go plugins for the protocol compiler

1- Install the protocol compiler plugins for Go using the following commands

$ export GO111MODULE=on  # Enable module mode
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

2- Update your PATH so that the protoc compiler can find the plugins 

$ export PATH="$PATH:$(go env GOPATH)/bin"

## Regenerate GRPC code

$ go install google.golang.org/protobuf/cmd/protoc-gen-go

$ mkdir ntpserver

$ protoc --go_out=plugins=grpc:ntpserver ntpserver.proto

## Compiling and executing 

### The server code
$ go run server/main.go

### The client code from another terminal
$ go run client/main.go

