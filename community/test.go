package main

import (
	pb "./protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:5000"
)

func runTest() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewRPCClient(conn)

	r, err := c.Apply(context.Background(), &pb.ApplyRequest{Applicant: "mars", Usercode: "sal"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Key: %s", r.Key)
}
