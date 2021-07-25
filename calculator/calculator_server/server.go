package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/rendaman0215/grpc/calculator/calculatorpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type server struct {
	calculatorpb.UnimplementedCalculateServiceServer
}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorReponse, error) {
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

func (*server) GetPrime(req *calculatorpb.GetPrimeRequest, stream calculatorpb.CalculateService_GetPrimeServer) error {
	input := req.GetInput()
	var k int32 = 2
	for input > 1 {
		if input%k == 0 {
			fmt.Printf("Devide %d by %d \n", input, k)
			res := &calculatorpb.GetPrimeResponse{
				Output: k,
			}
			stream.Send(res)
			input = input / k
		} else {
			k += 1
		}
	}
	return nil
}

func (*server) GetAvg(stream calculatorpb.CalculateService_GetAvgServer) error {
	var result int32 = 0
	var count int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := float32(result) / float32(count)
			return stream.SendAndClose(&calculatorpb.GetAvgResponse{
				Output: avg,
			})
		}
		if err != nil {
			log.Fatalf("Error while receiving client streaming: %v", err)
		}
		result = result + req.GetInput()
		count++
	}
}

func (*server) GetMax(stream calculatorpb.CalculateService_GetMaxServer) error {
	var max int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving client streaming: %v", err)
			return err
		}
		if req.GetInput() > max {
			max = req.GetInput()
			SendErr := stream.Send(&calculatorpb.GetMaxResponse{
				Output: max,
			})
			if SendErr != nil {
				log.Fatalf("Error while sending result: %v", SendErr)
			}
		}
	}
}

func (*server) SquareRoot(ctc context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Println("Received SquareRoot RPC...")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
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
