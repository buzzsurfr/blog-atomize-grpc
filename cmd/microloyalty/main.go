// test: grpcurl -proto proto/loyalty.proto -plaintext -d '{"customerID": 1234}' localhost:50051 loyalty.Points.GetPointsForCustomer
//
//go:generate protoc --plugin=/Users/tsalvo/go/bin/protoc-gen-go --plugin=/Users/tsalvo/go/bin/protoc-gen-go-grpc -I=../../proto --go_out=../../loyalty --go_opt=paths=source_relative --go-grpc_out=../../loyalty --go-grpc_opt=paths=source_relative ../../proto/loyalty.proto
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/buzzsurfr/blog-atomize-grpc/loyalty"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPointsServer
}

func (s *server) GetPointsForCustomer(ctx context.Context, in *pb.GetPointsRequest) (*pb.GetPointsReply, error) {
	// Mock that any customer has 100 loyalty points.
	loyaltyPoints := 100
	fmt.Printf("Loyalty Microservice: Customer ID %d has %d loyalty points.\n", in.CustomerID, loyaltyPoints)

	return &pb.GetPointsReply{
		CustomerID: in.CustomerID,
		Points:     int32(loyaltyPoints),
	}, nil
}

func (s *server) AddPoints(ctx context.Context, in *pb.AddPointsRequest) (*pb.AddPointsReply, error) {
	fmt.Printf("Loyalty Microservice: Added %d loyalty points to customer ID %d\n", in.Points, in.CustomerID)

	return &pb.AddPointsReply{}, nil
}

func main() {
	// Start Loyalty server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPointsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
