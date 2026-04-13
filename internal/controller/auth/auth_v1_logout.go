package auth

import (
	"context"

	v1 "backend/api/auth/v1"
	"backend/utility"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	utility.ClearAuthCookies(g.RequestFromCtx(ctx))
	return &v1.LogoutRes{}, nil
}
