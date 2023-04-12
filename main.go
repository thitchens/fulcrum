package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afeeblechild/fulcrum/lib"
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
	var err error
	logger, err = buildLogger()
	if err != nil {
		panic(err)
	}
}

func buildLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"fulcrum.log", "stderr"}
	return cfg.Build()
}

func loadConfig() (*Configuration, error) {
	file, err := os.Open("config.yaml")

	if err != nil {
		return nil, fmt.Errorf("cannot open config file: %v", err)
	}

	decoder := yaml.NewDecoder(file)
	config = &Configuration{}
	err = decoder.Decode(config)

	if err != nil {
		return nil, fmt.Errorf("cannot get configuration from file: %v", err)
	}

	return config, err
}
