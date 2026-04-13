package v1

import "github.com/gogf/gf/v2/frame/g"

type TeacherInfo struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

// POST /api/v1/auth/google/login

type GoogleLoginReq struct {
	g.Meta  `path:"/api/v1/auth/google/login" tags:"Auth" method:"post" summary:"Google OAuth login"`
	IdToken string `json:"idToken" v:"required"`
}

type GoogleLoginRes struct {
	Teacher TeacherInfo `json:"teacher"`
}

// GET /api/v1/auth/me

type MeReq struct {
	g.Meta `path:"/api/v1/auth/me" tags:"Auth" method:"get" summary:"Get current teacher profile"`
}

type MeRes struct {
	Teacher TeacherInfo `json:"teacher"`
}

// POST /api/v1/auth/logout

type LogoutReq struct {
	g.Meta `path:"/api/v1/auth/logout" tags:"Auth" method:"post" summary:"Logout"`
}

type LogoutRes struct{}
