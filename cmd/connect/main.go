package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"context"
	"github.com/glennsteven/kanban/internal/task"
	taskpb "github.com/glennsteven/proto/go/example/kanban/task/v1"
	"github.com/glennsteven/proto/go/example/kanban/task/v1/taskpbconnect"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	taskService := &task.Service{}
	path, handler := taskpbconnect.NewTaskServiceHandler(&taskConnectAdapter{service: taskService}, connect.WithInterceptors(validate.NewInterceptor()))
	mux.Handle(path, handler)
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)
	s := &http.Server{
		Addr:      "localhost:9998",
		Handler:   mux,
		Protocols: p,
	}
	log.Println("server started")
	return s.ListenAndServe()
}

type taskConnectAdapter struct {
	service *task.Service
}

func (a *taskConnectAdapter) Create(ctx context.Context, req *connect.Request[taskpb.CreateRequest]) (*connect.Response[taskpb.CreateResponse], error) {
	resp, err := a.service.Create(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	log.Printf("response %+v", resp)
	return connect.NewResponse(resp), nil
}
