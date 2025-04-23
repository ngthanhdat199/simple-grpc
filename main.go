package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"simple-grpc/gen/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreetServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	name := req.GetName()
	message := "Hello, " + name
	return &proto.HelloResponse{Message: message}, nil
}

func (s *server) ListConnections(ctx context.Context, req *proto.ListConnectionsReq) (*proto.ListConnectionsResp, error) {
	data := &proto.ListConnectionsResp{
		Data: &proto.Data{
			Connections: []*proto.Connection{
				{
					City:     "HoChiMinh",
					Country:  "Vietnam",
					Lat:      10.8231,
					Lon:      106.6297,
					AdminUrl: "https://cb-test3.connectbridge.net:443",
					MgmtUrl:  "https://cb-test3.connectbridge.net:443",
					Region:   "ap-southeast-1",
					FlagUrl:  "https://upload.wikimedia.org/wikipedia/commons/thumb/2/21/Flag_of_Vietnam.svg/1599px-Flag_of_Vietnam.svg.png?20170626140925",
				},
				{
					City:     "Germany",
					Country:  "Germany",
					Lat:      50.1109,
					Lon:      8.6821,
					AdminUrl: "https://cb-test4.connectbridge.net:443",
					MgmtUrl:  "https://cb-test4.connectbridge.net:443",
					Region:   "ap-southeast-1",
					FlagUrl:  "https://upload.wikimedia.org/wikipedia/commons/thumb/b/ba/Flag_of_Germany.svg/1600px-Flag_of_Germany.svg.png?20070926182838",
				},
			},
		},
	}

	return data, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreetServer(s, &server{})
	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
