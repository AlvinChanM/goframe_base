package middleware

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"go.opentelemetry.io/otel/trace"
)

// CustomLogger 自定义日志结构体
type CustomLogger struct{}

// NewLogger 创建新的日志实例
func NewLogger() *CustomLogger {
	return &CustomLogger{}
}

// WithFields 添加字段到日志
func (l *CustomLogger) WithFields(fields map[string]interface{}) *CustomLogger {
	// Implementation needed
	return l
}

// Info 记录info级别日志
func (l *CustomLogger) Info(msg string, kv ...interface{}) {
	l.log("INFO", msg, kv...)
}

// Error 记录error级别日志
func (l *CustomLogger) Error(msg string, kv ...interface{}) {
	l.log("ERROR", msg, kv...)
}

// Warning 记录warning级别日志
func (l *CustomLogger) Warning(msg string, kv ...interface{}) {
	l.log("WARNING", msg, kv...)
}

// Debug 记录debug级别日志
func (l *CustomLogger) Debug(msg string, kv ...interface{}) {
	l.log("DEBUG", msg, kv...)
}

// log 内部日志处理方法
func (l *CustomLogger) log(level string, msg string, kv ...interface{}) {
	// 构建完整的日志内容
	content := []interface{}{msg}
	if len(kv) > 0 {
		content = append(content, kv...)
	}

	// 获取上下文
	ctx := context.Background()

	// 根据日志级别调用相应的方法
	switch level {
	case "INFO":
		g.Log().Info(ctx, content...)
	case "ERROR":
		g.Log().Error(ctx, content...)
	case "WARNING":
		g.Log().Warning(ctx, content...)
	case "DEBUG":
		g.Log().Debug(ctx, content...)
	}
}

// StartSpan 创建新的span
func StartSpan(ctx context.Context, name string) trace.Span {
	_, span := gtrace.NewSpan(ctx, name)
	return span
}

// Example usage:
/*
func YourFunction(ctx context.Context) {
	// 创建新的span
	span := StartSpan(ctx, "your-operation")
	defer span.End()

	// 使用时创建logger
	logger := NewLogger()
	logger.Info(ctx, "开始操作")

	// 创建子span
	childSpan := StartSpan(ctx, "child-operation")
	defer childSpan.End()

	// 子操作的日志
	childLogger := NewLogger()
	childLogger.Info(ctx, "执行子操作")
}
*/
