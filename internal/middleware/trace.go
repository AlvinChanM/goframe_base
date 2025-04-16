package middleware

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

const (
	TraceIDKey = "trace_id"
)

// Trace 中间件用于注入追踪信息
func Trace(r *ghttp.Request) {
	// 生成唯一的 traceId
	traceId := guid.S()

	// 将 traceId 注入到上下文中
	ctx := context.WithValue(r.Context(), TraceIDKey, traceId)
	r.SetCtx(ctx)

	// 设置响应头
	r.Response.Header().Set("X-Trace-Id", traceId)

	// 继续处理请求
	r.Middleware.Next()
}

// GetTraceID 从上下文中获取 traceId
func GetTraceID(ctx context.Context) string {
	value := ctx.Value(TraceIDKey)
	if value == nil {
		return ""
	}
	if traceId, ok := value.(string); ok {
		return traceId
	}
	return ""
}
