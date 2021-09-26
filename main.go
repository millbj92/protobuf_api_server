package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"protobuf_api_example/ent"
	"protobuf_api_example/ent/migrate"
	"protobuf_api_example/ent/proto/entpb"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file")
    }

    client, err := ent.Open(
    dialect.MySQL, 
    fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/test", 
    viper.GetString("DB_USER"),
    viper.GetString("DB_PASSWORD"), 
    viper.GetString("DB_URL"), 
    viper.GetString("DB_PORT"),
    ))

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