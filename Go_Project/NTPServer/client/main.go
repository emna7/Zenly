
package main

import (
	"os"
	"time"
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "ntpserver/ntpserver"
)

const (
	address     = "localhost:9000"
	defaultName = "time.apple.com"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNtpServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetServer(ctx, &pb.Request{Ntpip: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Hello!: %s", r.Ntpresponse)
}
