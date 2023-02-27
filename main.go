package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/afeeblechild/fulcrum/lib"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func main() {
	lib.P("iam-manager", lib.Version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

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
	//TODO add check for logging functions to check if the file still exists
	file, err := os.OpenFile("fulcrum.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime)
}

func loadConfig() {
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}

	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		// TODO update panic
		panic(err)
	}
	err = t.Execute(w, r)
	if err != nil {
		panic(err)
	}
}
