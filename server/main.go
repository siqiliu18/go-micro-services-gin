package main

import (
	"context"
	"fmt"
	"log"
	"net"

	cs "computeserver/proto"

	"google.golang.org/grpc"
)

// Server is a type of struct
type Server struct{}

// Compute finds the GCD of 2 positive values.
func (s *Server) Compute(ctx context.Context, in *cs.GcdRequest) (*cs.GcdResponse, error) {
	a, b := in.A, in.B
	for b != 0 {
		a, b = b, a%b
	}
	out := cs.GcdResponse{
		Result: a,
	}
	return &out, nil
}

func main() {
	fmt.Println("Go gRPC server launching...")

	lis, err := net.Listen("tcp", ":7001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	svc := Server{}

	grpcServer := grpc.NewServer()

	cs.RegisterComputeServiceServer(grpcServer, &svc)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
