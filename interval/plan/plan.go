package plan

import (
	"context"
	"t-mk-opentrace/api/proto/plan"
	pm "t-mk-opentrace/api/proto/plan"
	"t-mk-opentrace/api/proto/task"
	"t-mk-opentrace/ext/grpc-driver/grpc"
	"t-mk-opentrace/ext/log-driver/log"
	rpc "t-mk-opentrace/services/rpc"
)

// RPCPlan RPCPlan
type RPCPlan struct {
}

// Search Search
func (p *RPCPlan) Search(ctx context.Context, res *plan.Request) (*plan.Response, error) {
	cc, err := rpc.NewTaskDial()
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
	return c.Delete(ct, &plan.Request{
		PlanID: res.PlanID,
	})
}

// Delete Search
func (p *RPCPlan) Delete(ctx context.Context, res *plan.Request) (*plan.Response, error) {
	_, _ = delTask(res.PlanID)
	_, _ = delTask(res.PlanID)
	return delTask(res.PlanID)
}

// delTask delTask
func delTask(planID string) (*plan.Response, error) {
	cc, err := rpc.NewTaskDial()
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
	c := task.NewTaskServiceClient(cc)
	resp, err := c.Search(ct, &task.TaskRequest{
		NameReq: planID,
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return &plan.Response{
		PlanName:  resp.NameRes,
		StartTime: resp.LocalAndServerTime,
		EndTime:   resp.LocalAndServerTime + "-end",
		Status:    pm.Response_Expires,
		Code:      pm.Response_success,
	}, nil
}
