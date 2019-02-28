package main

import (
	"../models"
	"./pb"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":5000"

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func runServer() {
	min := flag.Int("min", 5, "min person number")
	flag.Parse()

	pool := models.Pool{}
	pool.Init(*min)
	total, free := pool.GetStatus()
	fmt.Printf("total:%v free:%v\n", total, free)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		s := grpc.NewServer()
		pb.RegisterRpcServerServer(s, &server{})
		fmt.Printf("server listen on %v\n", port)
		_ = s.Serve(lis)
	}

}
