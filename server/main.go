package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

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
	Products map[string]*pb.Product
}

func printStructJSON(input interface{}) {
	val, _ := json.MarshalIndent(input, "", "  ")
	fmt.Println(string(val))
}

func newServer() *server {
	s := &server{
		Products: map[string]*pb.Product{
			"1305e1b4-bb31-4a18-9f06-261750d92beb": {
				Id:          "b770734e-5f86-4ea8-910a-e1a7a81a1926",
				Name:        "Minecraft",
				Description: "Just a video game",
				Picture:     "https://i.pinimg.com/564x/06/69/ef/0669efec5f8ac2ede45fa28a49fbaaba.jpg",
			},
			"9bc62ee1-2bf9-4cc7-b81d-71b3140815c0": {
				Id:          "55e002d3-d44e-4147-83e9-eab904e89060",
				Name:        "GAN Rubik",
				Description: "Best Rubik brand in the world",
				Picture:     "https://i.pinimg.com/564x/52/35/db/5235db65c46eb58c1f79a0d078f43fce.jpg",
			},
			"ff9e20f0-afa6-4618-8a07-2f4b2e894cd1": {
				Id:          "36f41028-8ee1-4507-81b9-73a46f3bf6f5",
				Name:        "Pad Thai",
				Description: "Thai's street food",
				Picture:     "https://i.pinimg.com/564x/eb/a8/c1/eba8c154a31061bf7ff989a45bd9727d.jpg",
			},
			"28f44977-e213-4351-a2e0-c3fd8a5be3df": {
				Id:          "677f7334-59e9-4001-8887-34a41e49d444",
				Name:        "Tom yum kung",
				Description: "Thai's street food",
				Picture:     "https://i.pinimg.com/564x/eb/a8/c1/eba8c154a31061bf7ff989a45bd9727d.jpg",
			},
			"8f35264a-61f1-4451-a85b-8d53670ed730": {
				Id:          "6da756cf-ba0d-43bb-853c-63bc4a6e967b",
				Name:        "Mug",
				Description: "Coffee glasses",
				Picture:     "https://i.pinimg.com/564x/eb/a8/c1/eba8c154a31061bf7ff989a45bd9727d.jpg",
			},
		}}
	return s
}

func (s *server) GetProduct(ctx context.Context, in *pb.Order) (*pb.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	if s.Products[in.Id] == nil {
		msg := fmt.Sprintf("error, product: %v not found", in.Id)
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}
	return s.Products[in.Id], nil
}

func (s *server) StreamProduct(in *pb.OrderArray, stream pb.Transfer_StreamProductServer) error {
	for i := range in.Id {
		if s.Products[in.Id[i]] != nil {
			if err := stream.Send(s.Products[in.Id[i]]); err != nil {
				log.Println(err.Error())
				return fmt.Errorf(err.Error())
			}
			fmt.Printf("product_id: %s has been streamed\n", s.Products[in.Id[i]].Id)
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

func (s *server) StreamOrder(stream pb.Transfer_StreamOrderServer) error {
	for {
		orderId, err := stream.Recv()
		if err == io.EOF {
			log.Println("stream success!")
			break
		}
		if err != nil {
			return err
		}
		if s.Products[orderId.Id] != nil {
			printStructJSON(s.Products[orderId.Id])
		} else {
			msg := fmt.Sprintf("error, product: %v not found", orderId.Id)
			log.Println(msg)
			return fmt.Errorf(msg)
		}
	}
	return nil
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
	pb.RegisterTransferServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
