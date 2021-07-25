package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/rendaman0215/grpc/greet/greetpb"
	grpc "google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	// サーバとコネクションを確立
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connet: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	//doUnary(c)

	//doServerStreaming(c)

	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kazuya",
			LastName:  "Saso",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManayTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kazuya",
			LastName:  "Saso",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Akashi",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Aomine",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Midorima",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Murasakibara",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Kise",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	// リクエストスライスから取り出して処理を回す
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}
