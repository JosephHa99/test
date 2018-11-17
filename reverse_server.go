package main

//THIS IS MAIN SERVER
import (
	"log"
	"net"

	pb "github.com/JosephHa99/test/msgreverse"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//server is used to implement reverse.server.
type server struct{}

func (s *server) SendInteger(ctx context.Context, in *pb.RequestInteger) (*pb.ReplyInteger, error) {
	var temp int32
	temp = in.Value * 2
	return &pb.ReplyInteger{Value: temp}, nil
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
	err := aServer.Serve(listener)
	//Check for errors
	if errorObj != nil {
		log.Fatalf("Failed to launch server: %v", err)
	}
}
