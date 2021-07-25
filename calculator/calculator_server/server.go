package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rendaman0215/grpc/calculator/calculatorpb"
	grpc "google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculateServiceServer
}

func (s *server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorReponse, error) {
	fmt.Printf("Start calculate...\n")
	fmt.Printf("Params: %v\n", req)
	First := req.GetCalculate().GetFirst()
	Second := req.GetCalculate().GetSecond()
	result := First + Second
	res := &calculatorpb.CalculatorReponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Server starts...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
