package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type User interface {
	Auth(ctx context.Context, user *UserAuthDTO) (*string, error)
	GetUser(ctx context.Context, user string) (*UserModel, error)
	GetAllUsers(ctx context.Context) (*[]UserModel, error)
}

type user struct {
}

func UserService(ctx context.Context) User {
	return &user{}
}

var testAuth UserAuthDTO

func (u *user) Auth(ctx context.Context, user *UserAuthDTO) (*string, error) {

	testAuth = UserAuthDTO{Login: "test", Password: "test"}

	if user.Login == testAuth.Login && user.Password == testAuth.Password {

		token, _ := getJwtToken("", "", "")

		return &token
	}

	return nil, errors.New("логин или пароль не корректны")

}

func (u *user) GetUser(ctx context.Context, user string) (*UserModel, error) {
	return &UserModel{
		ID:       uuid.New(),
		Login:    "test",
		Password: "pswd",
		Email:    "v@vendeta.ru",
		IsAdmin:  false,
	}, nil
}

func (u *user) GetAllUsers(ctx context.Context) (*[]UserModel, error) {
	return nil, nil
}
