package middleware

import (
	"strings"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// AccessLog 访问日志中间件
func AccessLog(r *ghttp.Request) {
	startTime := time.Now()
	logger := NewLogger()

	// 记录请求日志
	logger.Info("请求开始",
		"method", r.Method,
		"path", r.URL.Path,
		"client_ip", r.GetClientIp(),
		"user_agent", r.UserAgent(),
		"query", r.URL.RawQuery,
	)

	// 如果不是文件上传，记录请求体
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "multipart/form-data") {
		if body := r.GetBody(); len(body) > 0 {
			logger.Debug("请求体信息", "body", string(body))
		}
	}

	// 继续处理请求
	r.Middleware.Next()

	// 获取响应内容
	responseBody := r.Response.BufferString()
	var responseData interface{} = responseBody
	if len(responseBody) > 0 {
		// 尝试解析为JSON
		var responseMap map[string]interface{}
		if err := gconv.Struct(responseBody, &responseMap); err == nil {
			responseData = responseMap
		}
	}

	// 根据状态码决定日志级别
	logFunc := logger.Info
	if r.Response.Status >= 400 {
		logFunc = logger.Error
	}

	// 记录响应日志
	logFunc("请求结束",
		"status", r.Response.Status,
		"cost_seconds", time.Since(startTime).Seconds(),
		"latency", time.Since(startTime).String(),
		"response", responseData,
		"error", r.GetError(),
	)
}
