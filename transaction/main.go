package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/Zyko0/MonewayChallenge/transaction/pb"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) StoreTransaction(ctx context.Context, request *pb.TransactionRequest) (*pb.TransactionReply, error) {
	fmt.Println("content : " + request.String())
	rep := &pb.TransactionReply{Completed:true}
	return rep, nil
}
