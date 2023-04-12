package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type (
	DBConfiguration struct {
		DbName   string `yaml:"Dbname"`
		Password string `yaml:"Password"`
		Username string `yaml:"Username"`
	}
)

var (
	dbconfig *DBConfiguration
	Db       *sql.DB
)

func Init() error {
	err := loadConfig()
	if err != nil {
		return err
	}
	connect := fmt.Sprintf("dbname=%s sslmode=disable user=%s password=%s", dbconfig.DbName, dbconfig.Username, dbconfig.Password)
	Db, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func loadConfig() error {
	file, err := os.Open("db.yaml")

	if err != nil {
		return fmt.Errorf("cannot open db config file: %v", err)
	}

	decoder := yaml.NewDecoder(file)
	dbconfig = &DBConfiguration{}
	err = decoder.Decode(dbconfig)

	if err != nil {
		return fmt.Errorf("cannot get configuration from file: %v", err)
	}

	return err
}
