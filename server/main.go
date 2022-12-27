package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/Rayato159/go-gRPC-pg/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type config struct {
	Host *string
	Port *string
}

type server struct {
	pb.UnimplementedTransferServer
}

func (s *server) SendData(ctx context.Context, in *pb.Order) (*pb.Product, error) {
	products := map[string]*pb.Product{
		"1305e1b4-bb31-4a18-9f06-261750d92beb": {
			Id:          "b770734e-5f86-4ea8-910a-e1a7a81a1926",
			Name:        "Minecraft",
			Description: "Just a video game",
			Picture:     "https://i.pinimg.com/564x/06/69/ef/0669efec5f8ac2ede45fa28a49fbaaba.jpg",
		},
		"9bc62ee1-2bf9-4cc7-b81d-71b3140815c0": {
			Id:          "55e002d3-d44e-4147-83e9-eab904e89060",
			Name:        "GAN Rubik",
			Description: "Best Rubik brand of the world",
			Picture:     "https://i.pinimg.com/564x/52/35/db/5235db65c46eb58c1f79a0d078f43fce.jpg",
		},
		"ff9e20f0-afa6-4618-8a07-2f4b2e894cd1": {
			Id:          "36f41028-8ee1-4507-81b9-73a46f3bf6f5",
			Name:        "Pad Thai",
			Description: "Thai's street food",
			Picture:     "https://i.pinimg.com/564x/eb/a8/c1/eba8c154a31061bf7ff989a45bd9727d.jpg",
		},
	}
	if products[in.Id] == nil {
		msg := fmt.Sprintf("error, product: %v not found", in.Id)
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}
	return products[in.Id], nil
}

func main() {
	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error, can't load dotenv with an error: %v", err)
	}
	cfg := config{
		Host: flag.String("host", os.Getenv("HOST"), "The server host"),
		Port: flag.String("port", os.Getenv("PORT"), "The server port"),
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *cfg.Host, *cfg.Port))
	if err != nil {
		log.Fatalf("error, failed to listen: %v", err)
	}
	log.Printf("success, server is starting on %v:%v", *cfg.Host, *cfg.Port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTransferServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
