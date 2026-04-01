package cmd

import (
	"context"

	"backend/internal/controller/hello"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func ErrorHandle(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		gErrCode := gerror.Code(err)
		if gErrCode == gcode.CodeValidationFailed {
			r.Response.Status = 400
			r.Response.WriteJsonExit(g.Map{
				"message": "格式校驗失敗",
				"error":   "格式有誤",
				"code":    400,
				"data":    "格式有誤",
			})
			return
		}
	}
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ErrorHandle)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCSP)
				group.Middleware(MiddlewareCORS)
				group.Bind(
					hello.NewV1(),
				)
				// group.Middleware(MiddlewareAuth)
				// group.Bind(
				// )
				// group.Middleware(MiddlewareAdmin)
				// group.Bind(
				// )
				// group.Middleware(MiddlewareSuperAdmin)
				// group.Bind(
				// )
			})
			s.Run()
			return nil
		},
	}
)
