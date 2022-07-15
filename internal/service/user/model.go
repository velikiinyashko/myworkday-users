package user

import "github.com/google/uuid"

type UserModel struct {
	ID       uuid.UUID `json:"-" bson:"_id,omitempty"`
	Login    string    `json:"login" bson:"login"`
	Password string    `json:"-" bson:"password"`
	Email    string    `json:"email" bson:"email"`
	Phone    string    `json:"phone" bson:"phone"`
	IsAdmin  bool      `json:"-" bson:"is_admin"`
}

type UserAuthDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
