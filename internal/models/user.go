package models

import "context"

type User struct {
	ID       int
	FullName string
	Email    string
	Password string
	RoleID   int
}

type UserRepository interface {
	CreateUser(c context.Context , user User)(error)
}