package utility

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte
var jwtIssuer string

type JwtPurpose string

const (
	AccessToken     JwtPurpose = "access"
	RefreshToken    JwtPurpose = "refresh"
	ResetPassword   JwtPurpose = "reset_password"
	RegisterAccount JwtPurpose = "register_account"
	PublicRoom      JwtPurpose = "public_room"
)

type JwtClaims struct {
	UserId  uint64     `json:"user_id"`
	Purpose JwtPurpose `json:"purpose"`
	jwt.RegisteredClaims
}

type RegisterJwtClaims struct {
	Purpose JwtPurpose `json:"purpose"`
	jwt.RegisteredClaims
}

type JoinRoomClaims struct {
	TeamID  uint64 `json:"team_id"`
	MatchID uint64 `json:"match_id"` // 假設你要推播時根據比賽 ID 分房
	jwt.RegisteredClaims
}

type PublicRoomClaims struct {
	Purpose   JwtPurpose `json:"purpose"`    // 用來辨識是 public_room
	RoomToken string     `json:"room_token"` // 房間的唯一代碼
	Player1ID uint64     `json:"player1_id"` // 建立者 ID
	jwt.RegisteredClaims
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	jwtIssuer = os.Getenv("JWT_ISSUER")
}

func GenerateAccessToken(ctx context.Context, userId uint64, purpose JwtPurpose, expiration time.Duration) (string, error) {
	claims := JwtClaims{
		UserId:  userId,
		Purpose: purpose,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtIssuer,
			Subject:   fmt.Sprintf("%d", userId),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string, expectedPurpose JwtPurpose) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// check issuer is valid or not
	if claims.Issuer != jwtIssuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	// check purpose is valid or not
	if claims.Purpose != expectedPurpose {
		return nil, fmt.Errorf("invalid token purpose")
	}

	return claims, nil
}

func setCookieHeader(r *ghttp.Request, name, value string, maxAge int, path string, httpOnly bool) {

	secure := gcfg.Instance().GetBool("jwt.secureCookie") // read `config.yaml`
	domain := gcfg.Instance().GetString("jwt.domain")

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		HttpOnly: httpOnly,
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
	}
	r.Response.Header().Add("Set-Cookie", cookie.String())
}

func SetAuthTokens(ctx context.Context, r *ghttp.Request, userId uint64) error {
	// gen Access Token (12 hr)
	atoken, err := GenerateAccessToken(ctx, userId, AccessToken, 12*time.Hour)
	if err != nil {
		return err
	}

	// gen Refresh Token（30 days）
	rtoken, err := GenerateAccessToken(ctx, userId, RefreshToken, 30*24*time.Hour)
	if err != nil {
		return err
	}

	// 設定 Access Token 到 Header
	setCookieHeader(r, "access_token", atoken, 12*60*60, "/", true)

	// 設定 Refresh Token 到 Header
	setCookieHeader(r, "refresh_token", rtoken, 30*24*60*60, "/", true)

	return nil
}

func ClearAuthCookies(r *ghttp.Request) {
	setCookieHeader(r, "access_token", "", -1, "/", true)
	setCookieHeader(r, "refresh_token", "", -1, "/", true)
}
