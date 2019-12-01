package task

import (
	"context"
	"t-mk-opentrace/ext/log-driver/log"
	"t-mk-opentrace/proto/task"
	"time"
)

// Tb  t
type Tb struct {
}

// Search s
func (t *Tb) Search(ctx context.Context, r *task.TaskRequest) (*task.TaskResponse, error) {
	log.Warn(r)
	return &task.TaskResponse{
		Response:           r.GetRequest() + "serverRes",
		NameRes:            r.GetNameReq() + "app",
		AgeRes:             r.GetAgeReq() + 200,
		LocalAndServerTime: time.Now().String(),
	}, nil
}
