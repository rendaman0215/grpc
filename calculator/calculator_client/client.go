package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/rendaman0215/grpc/calculator/calculatorpb"
	grpc "google.golang.org/grpc"
)

func main() {
	// サーバとコネクションを確立
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connet: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculateServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.CalculatorRequest{
		Calculate: &calculatorpb.Calculate{
			First:  3,
			Second: 10,
		},
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculate RPC: %v", err)
	}

	log.Printf("Result of Calculation: %v", res.Result)
}

func doServerStreaming(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &calculatorpb.GetPrimeRequest{
		Input: 120,
	}
	resStream, err := c.GetPrime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetPrime RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GetPrime: %v", msg.GetOutput())
	}
}

func doClientStreaming(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*calculatorpb.GetAvgRequest{
		{
			Input: 1,
		},
		{
			Input: 2,
		},
		{
			Input: 3,
		},
		{
			Input: 4,
		},
	}

	stream, err := c.GetAvg(context.Background())
	if err != nil {
		log.Fatalf("Error while caling GetAvg RPC: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response from GetAvg: %v", err)
	}
	fmt.Printf("Result from GetAVG: %v\n", res.GetOutput())
}
