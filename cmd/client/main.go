package main

import (
	"context"
	taskpb "github.com/glennsteven/proto/go/example/kanban/task/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"log/slog"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conn, err := grpc.NewClient("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	taskClient := taskpb.NewTaskServiceClient(conn)
	resp, err := taskClient.Create(context.Background(), &taskpb.CreateRequest{
		Name: new("New TODO"),
		Desc: new("Todo Description"),
	})
	if err != nil {
		return err
	}

	l := slog.Default()

	rs, err := protojson.Marshal(resp)
	if err != nil {
		return err
	}
	l.Info("response", "body", string(rs))
	return nil
}
