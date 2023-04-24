package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afeeblechild/fulcrum/lib"
	"github.com/afeeblechild/fulcrum/lib/db"
	"github.com/afeeblechild/fulcrum/lib/log"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Address      string `yaml:"Address"`
	ReadTimeout  int64  `yaml:"ReadTimeout"`
	WriteTimeout int64  `yaml:"WriteTimeout"`
	Static       string `yaml:"Static"`
}

var (
	config *Configuration
	logger *zap.Logger
)

func main() {
	defer logger.Sync()
	defer db.DatabasePool.Close()

	fmt.Println("fulcrum", lib.Version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Defined in route_main.go
	mux.HandleFunc("/", rootHandler)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func init() {
	loadConfig()
	err := log.Init()
	if err != nil {
		panic(err)
	}
	err = db.Init()
	if err != nil {
		panic(err)
	}
}

func loadConfig() error {
	file, err := os.Open("config.yaml")

	if err != nil {
		return fmt.Errorf("cannot open config file: %v", err)
	}

	decoder := yaml.NewDecoder(file)
	config = &Configuration{}
	err = decoder.Decode(config)

	if err != nil {
		return fmt.Errorf("cannot get configuration from file: %v", err)
	}

	return err
}
