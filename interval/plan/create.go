package plan

import (
	"context"
	"net/http"
	"t-mk-opentrace/api/proto/task"
	"t-mk-opentrace/ext/grpc-driver/grpc"
	"t-mk-opentrace/ext/http-driver/requests"
	"t-mk-opentrace/ext/log-driver/log"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// SearchPlan test open trace
func SearchPlan(ctx *gin.Context) {
	span := opentracing.SpanFromContext(ctx.Request.Context())
	if span == nil {
		ctx.String(http.StatusInternalServerError, "Span is nil\n")
		span = opentracing.StartSpan("sp-lsp")
	}
	uri := "http://localhost:6006/user"
	if err := requests.Get(ctx.Request.Context(), uri, nil, nil); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.String(http.StatusOK, "router oks")
}

// TaskDial dial
func TaskDial() (*grpc.ClientConn, error) {
	return grpc.Dial("192.168.253.73:6005")
}

// TestTask 测试rpc连接
func TestTask(ctx context.Context) {
	cc, err := TaskDial()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := cc.Close(); err != nil {
			log.Warn(err)
		}
	}()
	c := task.NewTaskServiceClient(cc)
	// ctx, cancel := grpc.DefaultContext()
	// defer cancel()
	res, err := c.PlanDetail(ctx, &task.PlanRequest{
		PlanID: "planID-test-1",
	})
	if err != nil {
		log.Warn(err)
		return
	}
	log.Info(res.PlanID, res.GetCode())
}
