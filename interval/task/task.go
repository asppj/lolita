package task

import (
	"context"
	"t-mk-opentrace/api/proto/plan"
	"t-mk-opentrace/api/proto/task"
	"t-mk-opentrace/ext/grpc-driver/grpc"
	"t-mk-opentrace/ext/log-driver/log"
	"t-mk-opentrace/services/rpc"
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

// PlanDetail s
func (t *Tb) PlanDetail(ctx context.Context, r *task.PlanRequest) (*task.PlanResponse, error) {
	log.Warn(r)
	cc, err := rpc.NewPlanDial()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cc.Close(); err != nil {
			log.Warn()
		}
	}()
	ct, cancel := grpc.DefaultContext()
	defer cancel()
	c := plan.NewServiceClient(cc)
	resp, err := c.Search(ct, &plan.Request{
		PlanID: r.PlanID,
	})
	_, _ = c.Search(ct, &plan.Request{
		PlanID: r.PlanID,
	})
	if err != nil {
		return nil, err
	}
	return &task.PlanResponse{
		PlanID: resp.PlanName,
		Code:   int32(1),
	}, nil
}
