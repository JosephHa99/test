package main

import (
	"log"
	"os"
	"time"

	pb "github.com/JosephHa99/test"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	//Set up a connection to the server.
	connection, errorObj := grpc.Dial(address, grpc.WithInsecure())
	if errorObj != nil {
		log.Fatalf("Did not connect: %v", errorObj)
	}
	defer connection.Close()
	aClient := pb.NewGreeterClient(connection)

	//Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	contextObj, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, errorObj := aClient.SendInteger(context_obj, &pb.requestInteger{value: 2})
	if err != nil {
		log.Fatalf("Could not print: %v", errorObj)
	}
	log.Printf(str(response.Message))
}
