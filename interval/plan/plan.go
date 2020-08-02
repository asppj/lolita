package plan

import (
	"context"

	"github.com/asppj/lolita/ext/log-driver/log"
	rpc "github.com/asppj/lolita/pkg/plan/rpc"
	pm "github.com/asppj/lolita/proto/plan"
	"github.com/asppj/lolita/proto/task"
)

// RPCPlan RPCPlan
type RPCPlan struct {
}

// Search Search
func (p *RPCPlan) Search(ctx context.Context, res *pm.Request) (*pm.Response, error) {
	cc, err := rpc.NewTaskDial()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cc.Close(); err != nil {
			log.Warn()
		}
	}()
	c := pm.NewServiceClient(cc)
	return c.Delete(ctx, &pm.Request{
		PlanID: res.PlanID,
	})
}

// Delete Search
func (p *RPCPlan) Delete(ctx context.Context, res *pm.Request) (*pm.Response, error) {
	return delTask(ctx, res.PlanID)
}

// delTask delTask
func delTask(ctx context.Context, planID string) (*pm.Response, error) {
	cc, err := rpc.NewTaskDial()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cc.Close(); err != nil {
			log.Warn()
		}
	}()
	c := task.NewTaskServiceClient(cc)
	resp, err := c.Search(ctx, &task.TaskRequest{
		NameReq: planID,
	})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return &pm.Response{
		PlanName:  resp.NameRes,
		StartTime: resp.LocalAndServerTime,
		EndTime:   resp.LocalAndServerTime + "-end",
		Status:    pm.Response_Expires,
		Code:      pm.Response_success,
	}, nil
}
