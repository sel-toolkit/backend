package cmd

import (
	"context"
	"net/http"

	"backend/internal/controller/auth"
	"backend/internal/controller/hello"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func MiddlewareResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	err := r.GetError()
	if err != nil {
		gErrCode := gerror.Code(err)
		httpStatus := http.StatusInternalServerError
		code := 500

		switch gErrCode {
		case gcode.CodeValidationFailed:
			httpStatus = http.StatusBadRequest
			code = 400
		case gcode.CodeNotAuthorized:
			httpStatus = http.StatusUnauthorized
			code = 401
		case gcode.CodeNotFound:
			httpStatus = http.StatusNotFound
			code = 404
		}

		r.Response.Status = httpStatus
		r.Response.WriteJsonExit(Response{
			Code:    code,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	r.Response.WriteJsonExit(Response{
		Code:    200,
		Message: "success",
		Data:    r.GetHandlerResponse(),
	})
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(MiddlewareResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCSP)
				group.Middleware(MiddlewareCORS)
				group.Bind(
					hello.NewV1(),
					auth.NewV1(),
				)
				group.Middleware(MiddlewareAuth)
				group.Bind(
					auth.NewV1Protected(),
				)
			})
			s.Run()
			return nil
		},
	}
)
