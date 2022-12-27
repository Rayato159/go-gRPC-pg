package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/Rayato159/go-gRPC-pg/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	Host *string
	Port *string
}

const (
	defaultId string = "9bc62ee1-2bf9-4cc7-b81d-71b3140815c0"
)

func printStructJSON(input interface{}) {
	val, _ := json.MarshalIndent(input, "", "  ")
	fmt.Println(string(val))
}

func main() {
	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error, can't load dotenv with an error: %v\n", err)
	}
	cfg := config{
		Host: flag.String("host", os.Getenv("HOST"), "The server host"),
		Port: flag.String("port", os.Getenv("PORT"), "The server port"),
	}
	url := fmt.Sprintf("%s:%s", *cfg.Host, *cfg.Port)

	// product id
	productId := flag.String("product_id", defaultId, "Product id")

	flag.Parse()

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error, failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTransferClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.SendData(ctx, &pb.Order{
		Id: *productId,
	})
	if err != nil {
		log.Fatalf("could not send data with an: %v", err)
	}
	printStructJSON(r)
}
