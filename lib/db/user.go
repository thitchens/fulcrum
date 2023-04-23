package db

import (
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
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	prepared, err := Db.Prepare(statement)
	if err != nil {
		log.Error(err.Error())
	}
	defer prepared.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = prepared.QueryRow(uuid.New(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
