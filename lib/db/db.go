package db

import (
	"context"
	"fmt"
	"os"

	"github.com/afeeblechild/fulcrum/lib/log"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

type (
	Database struct {
		DatabaseName string `yaml:"DatabaseName"`
		DatabaseType string `yaml:"DatabaseType"` //Supported types: postgres
		Endpoint     string `yaml:"Endpoint"`
		Password     string `yaml:"Password"`
		Port         string `yaml:"Port"`
		Username     string `yaml:"Username"`
		Url          string
	}
)

var (
	DatabaseConfig *Database
	DatabasePool   *pgxpool.Pool
)

func Init() error {
	err := loadConfig()
	if err != nil {
		return err
	}
	DatabaseConfig.SetUrl()

	// pgxpool.New() creates a concurrency safe connection pool to use
	DatabasePool, err = pgxpool.New(context.Background(), DatabaseConfig.Url)
	if err != nil {
		panic(err)
	}
	return err
}

func (d *Database) SetUrl() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	d.Url = "postgres://" + d.Username + ":" + d.Password + "@" + d.Endpoint + ":" + d.Port + "/" + d.DatabaseName
}

func loadConfig() error {
	file, err := os.Open("db.yaml")

	if err != nil {
		return fmt.Errorf("cannot open db config file: %v", err)
	}

	decoder := yaml.NewDecoder(file)
	DatabaseConfig = &Database{}
	err = decoder.Decode(DatabaseConfig)

	if err != nil {
		return fmt.Errorf("cannot get configuration from file: %v", err)
	}

	return err
}

func Encrypt(value string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
	}
	return string(hash)
}
