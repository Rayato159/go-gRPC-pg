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

Server
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

Method
```go
func (s *server) SendData(ctx context.Context, in *pb.Order) (*pb.Product, error) { return &pb.Product{}}
```

Client
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

<h2>Start the server</h2>

Server terminal
```bash
go run server/main.go
```

Client terminal
```bash
go run client/main.go [--flag]
```
