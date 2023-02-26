package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/index.html")
	t.Execute(w, r)
}
