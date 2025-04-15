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
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 首先设置配置文件
			v, err := g.Cfg().GetWithEnv(ctx, "ENV")
			if err != nil {
				panic(err)
			}
			config_file := fmt.Sprintf("config.%s.yaml", v.String())
			fmt.Println(config_file)
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(config_file)

			// 测试数据库连接
			if err := g.DB().PingMaster(); err != nil {
				fmt.Println(err)
			}
			server := g.Cfg().MustGet(ctx, "server").Map()
			fmt.Println(server)

			// 然后创建服务器实例
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
