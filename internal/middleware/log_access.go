package middleware

import (
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// LogAccess 中间件用于记录访问日志
func LogAccess(r *ghttp.Request) {
	// 开始时间
	start := time.Now()

	// 获取请求信息
	requestData := map[string]interface{}{
		"trace_id":  GetTraceID(r.Context()),
		"method":    r.Method,
		"path":      r.URL.Path,
		"client_ip": r.GetClientIp(),
		"header":    r.Header,
	}

	// 如果不是文件上传，记录请求体
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "multipart/form-data") {
		if body := r.GetBody(); len(body) > 0 {
			requestData["body"] = string(body)
		}
	}

	// 记录请求日志
	g.Log().Info(r.Context(), "请求开始", requestData)

	// 继续处理请求
	r.Middleware.Next()

	// 记录响应日志
	responseData := map[string]interface{}{
		"trace_id":     GetTraceID(r.Context()),
		"status":       r.Response.Status,
		"cost_seconds": time.Since(start).Seconds(),
	}

	// 获取响应内容
	responseBody := r.Response.BufferString()
	if len(responseBody) > 0 {
		// 尝试解析为JSON，如果失败则作为字符串处理
		var responseMap map[string]interface{}
		if err := gconv.Struct(responseBody, &responseMap); err != nil {
			responseData["response"] = responseBody
		} else {
			responseData["response"] = responseMap
		}
	}

	g.Log().Info(r.Context(), "请求结束", responseData)
}
