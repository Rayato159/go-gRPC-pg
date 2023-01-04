package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	Key  string
}

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

	flag.Parse()

	// altsTC := alts.NewClientCreds(alts.DefaultClientOptions())
	// conn, err := grpc.Dial(url, grpc.WithTransportCredentials(altsTC))
	// if err != nil {
	// 	log.Fatalf("error, failed to connect: %v", err)
	// }
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error, failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTransferClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	r, err := client.GetProduct(ctx, &pb.Order{
		Id: "1305e1b4-bb31-4a18-9f06-261750d92beb",
	})
	if err != nil {
		log.Fatalf("could not send data with an: %v", err)
	}
	printStructJSON(r)

	orderIds := &pb.OrderArray{
		Id: []string{
			"1305e1b4-bb31-4a18-9f06-261750d92beb",
			"9bc62ee1-2bf9-4cc7-b81d-71b3140815c0",
			"ff9e20f0-afa6-4618-8a07-2f4b2e894cd1",
		},
	}

	streamProduct, err := client.StreamProduct(ctx, orderIds)
	if err != nil {
		log.Fatalf("could not stream product with an error: %v", err)
	}
	for {
		products, err := streamProduct.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.StreamProduct(_) = _, %v", client, err)
		}
		printStructJSON(products)
	}

	orders := []*pb.Order{
		{
			Id: "1305e1b4-bb31-4a18-9f06-261750d92beb",
		},
		{
			Id: "9bc62ee1-2bf9-4cc7-b81d-71b3140815c0",
		},
		{
			Id: "ff9e20f0-afa6-4618-8a07-2f4b2e894cd1",
		},
	}

	streamOrder, err := client.StreamOrder(ctx)
	if err != nil {
		log.Fatalf("could not stream order with an error: %v", err)
	}
	for i := range orders {
		if err := streamOrder.Send(orders[i]); err != nil {
			log.Fatalf("%v.Send(%v) = %v", streamOrder, orders[i], err)
		}
		fmt.Printf("order: %v has been streamed\n", orders[i].Id)
		time.Sleep(time.Millisecond * 500)
	}
	streamOrder.CloseAndRecv()

	streamAll, err := client.StreamAll(ctx)
	waitc := make(chan struct{})
	go func() {
		for {
			product, err := streamAll.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			printStructJSON(product)
		}
	}()
	for i := range orders {
		if err := streamAll.Send(orders[i]); err != nil {
			log.Fatalf("failed to send a note: %v", err)
		}
	}
	streamAll.CloseSend()
	<-waitc
}
