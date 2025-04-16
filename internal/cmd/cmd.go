package cmd

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/AlvinChanM/goframe_base/internal/controller/hello"
	"github.com/AlvinChanM/goframe_base/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化自定义日志格式
			logger := middleware.NewLogger()

			// 获取环境变量
			var configEnv string
			v, err := g.Cfg().GetWithEnv(ctx, "ENV")
			if err == nil && v.String() != "" {
				configEnv = v.String()
			}
			config_file := fmt.Sprintf("config.%s.yaml", configEnv)
			logger.Info("config_file:", config_file)

			// 设置配置文件路径和名称
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath("manifest/config")
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(config_file)

			// 获取并打印数据库配置，确保配置正确加载
			dbConfig := g.Cfg().MustGet(ctx, "database")
			if dbConfig.IsEmpty() {
				logger.Error("数据库配置未找到")
				return fmt.Errorf("数据库配置未找到")
			}
			middleware.NewLogger().WithFields(dbConfig.Map()).Info("Database config loaded")

			// 测试数据库连接
			if err := g.DB().PingMaster(); err != nil {
				return err
			}
			logger.Info("数据库连接成功")

			// 创建服务器实例
			s := g.Server()

			// 注册全局中间件
			s.Use(middleware.AccessLog, middleware.Trace)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})

			// 启动服务器
			s.Run()
			return nil
		},
	}
)
