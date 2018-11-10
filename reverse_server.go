package main

import (
	"log"
	"net"

	pb "github.com/JosephHa99/test"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//server is used to implement reverse.server.
type server struct{}

func (s *server) SendInteger(ctx context.Context, in *pb.requestInteger) (*pb.replyInteger, error) {
	return &pb.integerReply{Integer: "Value " + (in.value)*2}, nil
}

func main() {
	listener, errorObj := net.Listen("tcp", port)
	if errorObj != nil {
		log.Fatalf("Failed to listen: %v", errorObj)
	}
	aServer := grpc.NewServer()
	pb.RegisterIntegerMessageServer(aServer, &server{})
	//Register reflection service on gRPC server.
	reflection.Register(aServer)
	//run server
	errorObj := aServer.Serve(listener)
	//Check for errors
	if errorObj != nil {
		log.Fatalf("Failed to launch server: %v", err_obj)
	}
}
