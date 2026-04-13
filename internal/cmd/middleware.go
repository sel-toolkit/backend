package cmd

import (
	"net/http"
	"strings"

	"backend/utility"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = g.Cfg().GetStrings("cors.allowDomains")

	origin := r.Header.Get("Origin")

	// 若無 origin (非跨域工具)，直接放行
	if origin == "" {
		r.Response.CORS(corsOptions)
		r.Middleware.Next()
		return
	}

	// 動態回應符合的 Origin
	for _, domain := range corsOptions.AllowDomain {
		if domain == origin {
			r.Response.Header().Set("Access-Control-Allow-Origin", origin)
			r.Response.Header().Set("Access-Control-Allow-Credentials", "true")
			r.Response.CORS(corsOptions)
			r.Middleware.Next()
			return
		}
	}

	// 拒絕非白名單來源
	r.Response.WriteStatusExit(http.StatusForbidden)
}

func MiddlewareAuth(r *ghttp.Request) {
	tokenStr := r.Cookie.Get("access_token").String()
	if tokenStr == "" {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "Unauthorized")
		return
	}
	claims, err := utility.ValidateJWT(tokenStr, utility.AccessToken)
	if err != nil || claims == nil {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "Unauthorized")
		return
	}
	r.SetCtxVar("teacherId", int64(claims.UserId))
	r.Middleware.Next()
}

// func MiddlewareAdmin(r *ghttp.Request) {
// 	userId := r.GetCtxVar("userId").Uint64()
// 	// check if user is admin
// 	userRole, err := utility.GetUserRoleByUserId(userId)
// 	if err != nil {
// 		r.Response.WriteStatusExit(http.StatusInternalServerError, "Get user role failed")
// 		r.ExitAll()
// 	}
// 	if userRole != "Admin" && userRole != "SuperAdmin" {
// 		r.Response.WriteStatusExit(http.StatusForbidden, "Forbidden")
// 		r.ExitAll()
// 	}
// 	r.Middleware.Next()
// }

// func MiddlewareSuperAdmin(r *ghttp.Request) {
// 	userId := r.GetCtxVar("userId").Uint64()
// 	// check if user is super admin
// 	userRole, err := utility.GetUserRoleByUserId(userId)
// 	if err != nil {
// 		r.Response.WriteStatusExit(http.StatusInternalServerError, "Get user role failed")
// 		r.ExitAll()
// 	}
// 	if userRole != "SuperAdmin" {
// 		r.Response.WriteStatusExit(http.StatusForbidden, "Forbidden")
// 		r.ExitAll()
// 	}
// 	r.Middleware.Next()
// }

func MiddlewareCSP(r *ghttp.Request) {
	allowDomain := g.Cfg().GetStrings("cors.allowDomains")
	r.Response.Header().Set("Content-Security-Policy",
		"default-src 'self'; "+
			"script-src 'self' https://cdn.jsdelivr.net; "+
			"style-src 'self' https://fonts.googleapis.com; "+
			"font-src 'self' https://fonts.gstatic.com; "+
			"img-src 'self' data:; "+
			"connect-src 'self' "+strings.Join(allowDomain, " ")+"; ")

	r.Middleware.Next()
}
