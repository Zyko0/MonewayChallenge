package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/Zyko0/MonewayChallenge/transaction/pb"
	"github.com/Zyko0/MonewayChallenge/balance/pb"
	"github.com/Zyko0/MonewayChallenge/db"
)

type server struct{}

var balanceClient balance.BalanceServiceClient

func main() {
	// Setup dial with balance service
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	balanceClient = balance.NewBalanceServiceClient(conn)

	// Setting up grpc server
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
	// Storing the transaction into the database
	db.StoreTransaction(request)
	// Credit or debit the balance by sending a new request to BalanceService
	balanceRequest := &balance.BalanceRequest{
		AccountID:request.AccountID,
		Value:request.Amount,
		Currency:request.Currency,
	}
	res, err := balanceClient.ManageBalance(ctx, balanceRequest)
	if err != nil || !res.Completed {
		rep.Completed = false
		return rep, err
	}
	return rep, nil
}
