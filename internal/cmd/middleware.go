package cmd

import (
	"net/http"
	"strings"

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

// func MiddlewareAuth(r *ghttp.Request) {
// 	aToken := r.Cookie.Get("access_token").String()
// 	rToken := r.Cookie.Get("refresh_token").String()
// 	aTokenClaims, err := utility.ValidateJWT(aToken, utility.AccessToken)
// 	if err != nil || aTokenClaims == nil {
// 		// refresh token
// 		rTokenClaims, err := utility.ValidateJWT(rToken, utility.RefreshToken)
// 		if err != nil || rTokenClaims == nil {
// 			r.Response.WriteStatusExit(http.StatusUnauthorized, "Unauthorized")
// 			r.ExitAll()
// 		}
// 		err = utility.SetAuthTokens(r.Context(), r, rTokenClaims.UserId)
// 		if err != nil {
// 			r.Response.WriteStatusExit(http.StatusInternalServerError, "Token refresh failed")
// 			r.ExitAll()
// 		}
// 		r.SetCtxVar("userId", rTokenClaims.UserId)
// 	} else {
// 		r.SetCtxVar("userId", aTokenClaims.UserId)
// 	}
// 	r.Middleware.Next()
// }

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
