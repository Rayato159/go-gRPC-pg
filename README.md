<h1>gRPC with Go</h1>

<h2>Build a proto</h2>

`.proto` example

```proto
syntax = "proto3"; // required

option go_package = "your-package-name"; // required

message MessageName {
    opt type field_name = order
    ...
}

service ServiceName {
    rpc func_name (request) returns (response)
    ...
}
```

Build go-proto command
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/shop.proto
```

<h2>Quick Start</h2>

<h3>Server</h3>

```go
//...
lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *cfg.Host, *cfg.Port))
if err != nil {
    log.Fatalf("error, failed to listen: %v", err)
}
log.Printf("success, server is starting on %v:%v", *cfg.Host, *cfg.Port)

var opts []grpc.ServerOption
grpcServer := grpc.NewServer(opts...)
pb.RegisterTransferServer(grpcServer, &server{})
grpcServer.Serve(lis)
```

<strong>Method</strong>

Simple RPC
```go
func (s *server) SendData(ctx context.Context, in *pb.Order) (*pb.Product, error) {
    return &pb.Product{}
}
```

<strong>Server-side streaming RPC</strong>

Sever -> `stream.Send()` is for stream a data
```go
func (s *server) StreamProduct(in *pb.OrderArray, stream pb.Transfer_StreamProductServer) error {
	for i := range in.Id {
        //...
		if s.Products[in.Id[i]] != nil {
			if err := stream.Send(s.Products[in.Id[i]]); err != nil {
                //...
			}
		}
	}
    //...
	return nil
}
```

<h3>Client</h3>

```go
//...
conn, err := grpc.Dial(addrUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
if err != nil {
    //...
}
defer conn.Close()
c := pb.NewTransferClient(conn)

ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
defer cancel()

r, err := c.SendData(ctx, &pb.Order{
    Id: *productId,
})
if err != nil {
    //...
}
```

For receive a streaming data from server `stream.Recv()` to Receive a data and `io.EOF` is mean the sever function has been returned
```go
stream, err := client.StreamProduct(ctx, orderIds)
if err != nil {
    //...
}
for {
    products, err := stream.Recv()
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatalf("%v.StreamProduct(_) = _, %v", client, err)
    }
    //...
}
```

<h2>Start the server</h2>

Server terminal
```bash
go run server/main.go
```

Client terminal
```bash
go run client/main.go [--flag]
```
