package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/rendaman0215/grpc/calculator/calculatorpb"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// doClientStreaming(c)
	// doBiDiStreaming(c)
	doErrorUnary(c)
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

func doBiDiStreaming(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting to do a Bi-Directional Streaming RPC...")

	requests := []*calculatorpb.GetMaxRequest{
		{
			Input: 1,
		},
		{
			Input: 5,
		},
		{
			Input: 3,
		},
		{
			Input: 6,
		},
		{
			Input: 2,
		},
		{
			Input: 20,
		},
	}

	stream, err := c.GetMax(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GetMax: %v", err)
	}

	waitc := make(chan struct{})
	// メッセージを任意の数送信 (go routine)
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending data: %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while sending data: %v", err)
			}
		}
		stream.CloseSend()
	}()
	// メッセージを任意の数受信 (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error thile receiving the data: %v", err)
			}
			fmt.Printf("The Maximum Number for now is: %v\n", res.GetOutput())
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// 正常系
	doErrorCall(c, 10)

	// 異常系
	doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculateServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
			return
		}
	}

	log.Printf("Result of SquareRoot: %v: %v \n", n, res.GetNumberRoot())
}
