package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yermanovberik/golang-final-project/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool)(models.UserRepository){
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(c context.Context , user models.User)(error){
	return nil
}