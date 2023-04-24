// test: grpcurl -proto proto/orders.proto -plaintext -d '{"customerID": 1234}' localhost:50051 loyalty.Points.GetPointsForCustomer
//
//go:generate protoc --plugin=/Users/tsalvo/go/bin/protoc-gen-go --plugin=/Users/tsalvo/go/bin/protoc-gen-go-grpc -I=../../proto --go_out=../../orders --go_opt=paths=source_relative --go-grpc_out=../../orders --go-grpc_opt=paths=source_relative ../../proto/orders.proto
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	lpb "github.com/buzzsurfr/blog-atomize-grpc/loyalty"
	pb "github.com/buzzsurfr/blog-atomize-grpc/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	loyaltyClient lpb.PointsClient
)

type server struct {
	pb.UnimplementedOrderServer
}

func (s *server) PlaceOrder(ctx context.Context, in *pb.PlaceOrderRequest) (*pb.PlaceOrderReply, error) {
	// Process the order itself
	fmt.Printf("Orders Microservice: Order placed by customer %d for %d cents.\n", in.CustomerID, in.Cents)

	// Add loyalty points based on the order's amount.
	// 1 LP per cent (USD)
	loyaltyPoints := in.Cents
	if _, err := loyaltyClient.AddPoints(context.Background(), &lpb.AddPointsRequest{CustomerID: in.CustomerID, Points: loyaltyPoints}); err != nil {
		fmt.Printf("Orders Microservice: Something happened while adding %d loyalty points for customer ID %d.\n", loyaltyPoints, in.CustomerID)
	} else {
		fmt.Printf("Orders Microservice: %d loyalty points added to customer with ID %d\n", loyaltyPoints, in.CustomerID)
	}

	// Assuming the order went through even if the points weren't assigned
	return &pb.PlaceOrderReply{}, nil
}

func main() {
	// Connect to Loyalty service
	loyaltyConn, loyaltyErr := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if loyaltyErr != nil {
		log.Fatalf("did not connect to loyalty: %v", loyaltyErr)
	}
	defer loyaltyConn.Close()
	loyaltyClient = lpb.NewPointsClient(loyaltyConn)

	// Start orders server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
