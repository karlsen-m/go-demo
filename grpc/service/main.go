package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "service/test" // 导入您的 gRPC 服务定义包
)

type server struct {
	pb.UnimplementedTestServer
}

func (s *server) GetTest(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	// 在这里实现您的具体业务逻辑
	response := &pb.Response{
		Message: "Hello, " + request.Name,
	}
	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})

	log.Println("gRPC server is running on port 5002")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
