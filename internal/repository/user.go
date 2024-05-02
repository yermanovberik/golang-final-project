package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yermanovberik/golang-final-project/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool)(models.UserRepository){
	return &UserRepository{db: db}
}


func (ur *UserRepository) CreateUser(c context.Context, user models.UserRequest) (int, error) {
	var userID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO public.users(
		full_name, email, password, role, created_at)
		VALUES ($1, $2, $3, $4, $5) returning id;`
	err := ur.db.QueryRow(c, userQuery, user.FullName, user.Email, user.Password,"user" , currentTime).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (ur *UserRepository) EditUser(c context.Context, user models.User) (int, error) {
	userQuery := `UPDATE users
	SET full_name=$1, email=$2 , role=$3
	WHERE id = $4`
	_, err := ur.db.Exec(c, userQuery, user.FullName,user.Email, user.Role , user.ID)
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil
}

func (ur *UserRepository) DeleteUser(c context.Context, userID int) error {
	query := `delete from users where id = $1`
	_, err := ur.db.Exec(c, query, userID)
	if err != nil {
		return err
	}
	return nil
}


func (ur *UserRepository) GetUserByEmail(c context.Context, email string) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email , fullname, password, role, created_at FROM users where email = $1`
	row := ur.db.QueryRow(c, query, email)
	err := row.Scan(&user.ID, &user.Email,&user.FullName, &user.Password, &user.Role , &user.CreatedAt)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByID(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email , fullname, password, role, created_at FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.ID, &user.Email,&user.FullName, &user.Password, &user.Role , &user.CreatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}