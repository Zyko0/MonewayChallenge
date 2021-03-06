package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/Zyko0/MonewayChallenge/balance/pb"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	balance.RegisterBalanceServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) ManageBalance(ctx context.Context, request *balance.BalanceRequest) (*balance.BalanceReply, error) {
	fmt.Println("content : " + request.String())
	rep := &balance.BalanceReply{Completed:true}
	// Storing the new balance into the database

	return rep, nil
}


