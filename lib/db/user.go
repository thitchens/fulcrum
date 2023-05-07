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

func (user *User) CreateSession() (Session, error) {
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at;"

	var session Session
	err := DatabasePool.QueryRow(context.TODO(), statement, uuid.New(), user.Email, user.Id, time.Now().UTC()).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return session, err
}

func (user *User) GetSession() (Session, error) {
	statement := "SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id=$1;"
	session := Session{}

	err := DatabasePool.QueryRow(context.TODO(), statement, user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	return session, err
}

func (user *User) Create() error {
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at;"

	return DatabasePool.QueryRow(context.TODO(), statement, uuid.New(), user.Name, user.Email, HashPassword(user.Password), time.Now().UTC()).
		Scan(&user.Id, &user.Uuid, &user.CreatedAt)
}

func (user *User) Update() error {
	// TODO not sure if I want to update all at once, or have separate functions for each field
	statement := "UPDATE users SET name=$2, email=$3, password=$4 WHERE id=$1;"

	_, err := DatabasePool.Exec(context.TODO(), statement, user.Name, user.Email, HashPassword(user.Password))
	return err
}

func (user *User) Delete() error {
	statement := "DELETE FROM users WHERE uuid=$1;"

	_, err := DatabasePool.Exec(context.TODO(), statement, user.Uuid)
	return err
}

func UserByEmail(email string) (*User, error) {
	user := User{}

	err := DatabasePool.QueryRow(context.TODO(), "SELECT id, uuid, name, email, password, created_at FROM users WHERE email=$1;", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return &user, err
}
