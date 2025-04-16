package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/AlvinChanM/goframe_base/api/hello/v1"
	"github.com/AlvinChanM/goframe_base/internal/middleware"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	traceId := middleware.GetTraceID(ctx)
	g.Log().Info(ctx, "处理请求", "trace_id", traceId)
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
