package main

import (
	"log"
	"time"

	pb "github.com/JosephHa99/test/msgreverse"
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
	aClient := pb.NewIntegerMessageClient(connection)

	//Contact the server and print out its response.
	/*
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
	*/
	contextObj, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, errorObj := aClient.SendInteger(contextObj, &pb.RequestInteger{Value: 140})
	if errorObj != nil {
		log.Fatalf("Could not print: %v", errorObj)
	}
	log.Printf("Integer: %d", response.Value)
	log.Printf("%d", response.Value)
}
