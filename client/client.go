package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"simple-grpc/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreetClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Gopher"})
	// if err != nil {
	// 	log.Fatalf("Error while calling SayHello: %v", err)
	// }

	resp, err := client.ListConnections(ctx, &proto.ListConnectionsReq{})
	if err != nil {
		log.Fatalf("Error while calling ListConnections: %v", err)
	}

	log.Printf("Response from server:")
	PrintJSON(resp)
}

func PrintJSON(obj interface{}) {
	prettyJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
