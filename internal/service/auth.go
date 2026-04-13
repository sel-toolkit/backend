package service

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/dao"
	"backend/internal/model/do"
	"backend/internal/model/entity"
	"backend/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/api/idtoken"
)

type sAuth struct{}

var Auth = &sAuth{}

// VerifyGoogleToken validates a Google ID token and returns its payload.
func (s *sAuth) VerifyGoogleToken(ctx context.Context, idToken string) (*idtoken.Payload, error) {
	clientID := g.Cfg().MustGet(ctx, "google.clientId").String()
	payload, err := idtoken.Validate(ctx, idToken, clientID)
	if err != nil {
		return nil, gerror.Wrap(err, "invalid google id token")
	}
	return payload, nil
}

// UpsertTeacher creates or updates a teacher record from a Google ID token payload.
func (s *sAuth) UpsertTeacher(ctx context.Context, payload *idtoken.Payload) (*entity.Teachers, error) {
	googleSub := payload.Subject
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	avatarUrl, _ := payload.Claims["picture"].(string)

	cols := dao.Teachers.Columns()

	var teacher entity.Teachers
	err := dao.Teachers.Ctx(ctx).Where(cols.GoogleSub, googleSub).Scan(&teacher)
	if err != nil && err != sql.ErrNoRows {
		return nil, gerror.Wrap(err, "query teacher failed")
	}

	if teacher.Id == 0 {
		id, err := dao.Teachers.Ctx(ctx).Data(do.Teachers{
			GoogleSub: googleSub,
			Email:     email,
			Name:      name,
			AvatarUrl: avatarUrl,
		}).InsertAndGetId()
		if err != nil {
			return nil, gerror.Wrap(err, "create teacher failed")
		}
		teacher.Id = id
		teacher.Email = email
		teacher.Name = name
		teacher.AvatarUrl = avatarUrl
		teacher.GoogleSub = googleSub
	} else {
		_, err = dao.Teachers.Ctx(ctx).Where(cols.Id, teacher.Id).Data(do.Teachers{
			Name:      name,
			AvatarUrl: avatarUrl,
		}).Update()
		if err != nil {
			return nil, gerror.Wrap(err, "update teacher failed")
		}
		teacher.Name = name
		teacher.AvatarUrl = avatarUrl
	}

	return &teacher, nil
}

// GenerateToken signs a JWT for the given teacher ID.
func (s *sAuth) GenerateToken(ctx context.Context, teacherId int64) (string, error) {
	token, err := utility.GenerateAccessToken(ctx, uint64(teacherId), utility.AccessToken, 24*time.Hour)
	if err != nil {
		return "", gerror.Wrap(err, "generate token failed")
	}
	return token, nil
}

// GetTeacherById retrieves a teacher by their primary key.
func (s *sAuth) GetTeacherById(ctx context.Context, teacherId int64) (*entity.Teachers, error) {
	var teacher entity.Teachers
	err := dao.Teachers.Ctx(ctx).Where(dao.Teachers.Columns().Id, teacherId).Scan(&teacher)
	if err != nil {
		return nil, gerror.Wrap(err, "query teacher failed")
	}
	if teacher.Id == 0 {
		return nil, gerror.New("teacher not found")
	}
	return &teacher, nil
}
