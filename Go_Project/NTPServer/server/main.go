package main

import (
    "net"

    "log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	ntp "github.com/beevik/ntp"
	pb "ntpserver/ntpserver"
)

// Server interface for our service methods
type server struct {
	pb.UnimplementedNtpServiceServer
}

// GetBook logs Book from client and returns new Book
func (s *server) GetServer(ctx context.Context, input *pb.Request) (*pb.Response, error) {

	time, err := ntp.Time(input.Ntpip)
	if (err == nil) {
		return &pb.Response{Ntpresponse: time.String()}, nil
	} else {
		return &pb.Response{Ntpresponse: "Unable to request from"}, nil
	}
}

func main() {

    var port = "9000";

    l, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Printf("Failed to listen")
    }

    s := grpc.NewServer()
	pb.RegisterNtpServiceServer(s, &server{})

    log.Printf("gRPC server started at 9000")
    if err := s.Serve(l); err != nil {
        log.Printf("Failed to launch server")
    }

}
