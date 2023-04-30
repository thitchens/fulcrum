package db

import (
	"context"
	"time"

	"github.com/afeeblechild/fulcrum/lib/log"
	"github.com/google/uuid"
)

type (
	User struct {
		Id        int
		Uuid      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}

	Session struct {
		Id        int
		Uuid      string
		Email     string
		UserId    int
		CreatedAt time.Time
	}
)

func (user *User) Create() error {
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"

	err := DatabasePool.QueryRow(context.TODO(), statement, uuid.New(), user.Name, user.Email, Encrypt(user.Password), time.Now().UTC()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}

func UserByEmail(email string) (*User, error) {
	statement := "SELECT id, uuid, name, email, password, created_at FROM users WHERE email='$1'"
	user := User{}

	err := DatabasePool.QueryRow(context.TODO(), statement, email).Scan(user.Id, user.Uuid, user.Name, user.Email, user.Password, user.CreatedAt)

	return &user, err
}