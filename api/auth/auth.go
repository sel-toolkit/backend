// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"backend/api/auth/v1"
)

// IAuthV1 covers public endpoints (no auth required).
type IAuthV1 interface {
	GoogleLogin(ctx context.Context, req *v1.GoogleLoginReq) (res *v1.GoogleLoginRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}

// IAuthV1Protected covers endpoints that require a valid JWT.
type IAuthV1Protected interface {
	Me(ctx context.Context, req *v1.MeReq) (res *v1.MeRes, err error)
}
