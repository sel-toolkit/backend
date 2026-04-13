package auth

import (
	"context"

	v1 "backend/api/auth/v1"
	"backend/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1Protected) Me(ctx context.Context, req *v1.MeReq) (res *v1.MeRes, err error) {
	teacherId := g.RequestFromCtx(ctx).GetCtxVar("teacherId").Int64()
	if teacherId == 0 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "unauthorized")
	}

	teacher, err := service.Auth.GetTeacherById(ctx, teacherId)
	if err != nil {
		return nil, err
	}

	return &v1.MeRes{
		Teacher: v1.TeacherInfo{
			Id:        teacher.Id,
			Email:     teacher.Email,
			Name:      teacher.Name,
			AvatarUrl: teacher.AvatarUrl,
		},
	}, nil
}
