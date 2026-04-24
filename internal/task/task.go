package task

import (
	"context"
	wktpb "github.com/glennsteven/proto/go/common/wkt/v1"
	taskpb "github.com/glennsteven/proto/go/example/kanban/task/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ taskpb.TaskServiceServer = new(Service)

type Service struct{}

func (s *Service) Create(ctx context.Context, req *taskpb.CreateRequest) (*taskpb.CreateResponse, error) {
	now := time.Now()
	t := taskpb.Task{
		Id:     new(uuid.NewString()),
		Name:   new(req.GetName()),
		Desc:   new(req.GetDesc()),
		Status: new(taskpb.Status_BACKLOG),
		Timestamps: &wktpb.RecordTimestamps{
			CreatedAt: timestamppb.New(now),
			UpdatedAt: timestamppb.New(now),
			DeletedAt: nil,
		},
	}
	resp := &taskpb.CreateResponse{
		Task: &t,
	}
	return resp, nil
}
