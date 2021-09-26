package main

import (
	"context"
	"log"
	"net"

	"protobuf_api_example/ent"
	"protobuf_api_example/ent/migrate"
	"protobuf_api_example/ent/proto/entpb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"entgo.io/ent/dialect"
)

func main() {
    client, err := ent.Open(dialect.MySQL, "root:Eq9935sm!@tcp(localhost:3306)/nuboverflow_users")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()

    ctx := context.Background()
    // Run migration.
    err = client.Schema.Create(
        ctx,
		migrate.WithGlobalUniqueID(true),
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true), 
    )

    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

     svc := entpb.NewUserService(client)

    // Create a new gRPC server (you can wire multiple services to a single server).
    server := grpc.NewServer()

    // Register the User service with the server.
    entpb.RegisterUserServiceServer(server, svc)

    // Open port 5000 for listening to traffic.
    lis, err := net.Listen("tcp", ":5000")
    if err != nil {
        log.Fatalf("failed listening: %s", err)
    }

    // Listen for traffic indefinitely.
    if err := server.Serve(lis); err != nil {
        log.Fatalf("server ended: %s", err)
    }

}