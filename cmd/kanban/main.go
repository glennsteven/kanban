package main

import (
	"github.com/glennsteven/kanban/internal/task"
	taskpb "github.com/glennsteven/proto/go/example/kanban/task/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listen, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		return err
	}
	defer listen.Close()

	srv := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	taskpb.RegisterTaskServiceServer(srv, &task.Service{})
	log.Printf("server listening at %v", listen.Addr())
	return srv.Serve(listen)
}
