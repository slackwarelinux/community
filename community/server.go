package main

import (
	"../models"
	"./protobuf"
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var pool = models.Pool{}

type server struct{}

func (s *server) Apply(ctx context.Context, in *protobuf.ApplyRequest) (*protobuf.NormalReply, error) {
	res := protobuf.NormalReply{}
	person, err := pool.Apply()
	if err == nil {
		person.Code = in.Usercode
		person.Applicant = in.Applicant
		res.Error = 0
		res.Msg = "OK"
		res.Key = person.Id
	} else {
		res.Error = -1
		res.Msg = "apply fial."
	}
	total, free := pool.GetStatus()
	log.Printf("Total:%v,Free:%v", total, free)
	return &res, nil
}

func runServer() {
	min := flag.Int("min", 5, "min person number")
	address := flag.String("address", ":5000", "listen address. default :5000")
	flag.Parse()

	pool.Init(*min)

	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		s := grpc.NewServer()
		protobuf.RegisterRPCServer(s, &server{})
		log.Printf("server listen on %v\n", *address)
		_ = s.Serve(lis)
	}

}
