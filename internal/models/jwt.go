package models

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Role string `json:"role"`
	ID   int   `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	Role string `json:"role"`
	ID int `json:"id"`
	jwt.RegisteredClaims
}