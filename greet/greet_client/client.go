package main

import (
	"fmt"
	"log"

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
	fmt.Printf("Created client %f", c)
}
