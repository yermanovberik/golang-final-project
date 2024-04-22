package models

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	RoleID int `json:"role"`
	ID   int   `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	RoleID int `json:"role"`
	ID int `json:"id"`
	jwt.RegisteredClaims
}