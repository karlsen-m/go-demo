package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "client/test" // 导入您的 gRPC 服务定义包
)

func main() {
	conn, err := grpc.Dial(":5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTestClient(conn)

	request := &pb.Request{
		Name: "John",
	}

	response, err := client.GetTest(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to call YourMethod: %v", err)
	}

	log.Println("Response:", response.Message)
}
