package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/velikiinyashko/myworkday/model"
)

type User interface {
	Auth(ctx context.Context, user *model.UserAuthDTO) error
	GetUser(ctx context.Context, user string) (*model.User, error)
	GetAllUsers(ctx context.Context) (*[]model.User, error)
}

type user struct {
}

func UserService(ctx context.Context) User {
	return &user{}
}

var testAuth model.UserAuthDTO

func (u *user) Auth(ctx context.Context, user *model.UserAuthDTO) error {

	testAuth = model.UserAuthDTO{Login: "test", Password: "test"}

	if user.Login == testAuth.Login && user.Password == testAuth.Password {
		return nil
	}

	return errors.New("логин или пароль не корректны")

}

func (u *user) GetUser(ctx context.Context, user string) (*model.User, error) {
	return &model.User{
		ID:       uuid.New(),
		Login:    "test",
		Password: "pswd",
		Email:    "v@vendeta.ru",
		IsAdmin:  false,
	}, nil
}

func (u *user) GetAllUsers(ctx context.Context) (*[]model.User, error) {
	return nil, nil
}
