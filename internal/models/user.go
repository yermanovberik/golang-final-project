package models

import "context"

type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Role      string
	CreatedAt string
}

type UserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, userID int) (User, error)

	CreateUser(c context.Context, user UserRequest) (int, error)
	EditUser(c context.Context, user User) (int, error)
	DeleteUser(c context.Context, userID int) error
}
