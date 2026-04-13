package auth

import (
	"context"

	v1 "backend/api/auth/v1"
	"backend/internal/service"
	"backend/utility"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GoogleLogin(ctx context.Context, req *v1.GoogleLoginReq) (res *v1.GoogleLoginRes, err error) {
	payload, err := service.Auth.VerifyGoogleToken(ctx, req.IdToken)
	if err != nil {
		return nil, err
	}

	teacher, err := service.Auth.UpsertTeacher(ctx, payload)
	if err != nil {
		return nil, err
	}

	r := g.RequestFromCtx(ctx)
	if err = utility.SetAuthTokens(ctx, r, uint64(teacher.Id)); err != nil {
		return nil, err
	}

	return &v1.GoogleLoginRes{
		Teacher: v1.TeacherInfo{
			Id:        teacher.Id,
			Email:     teacher.Email,
			Name:      teacher.Name,
			AvatarUrl: teacher.AvatarUrl,
		},
	}, nil
}
