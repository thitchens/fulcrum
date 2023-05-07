package main

import (
	"fmt"
	"net/http"

	"github.com/afeeblechild/fulcrum/lib/db"
	"github.com/afeeblechild/fulcrum/lib/log"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// rootTemplate := "web/templates/index.html"
	// t, err := template.ParseFiles(rootTemplate)
	// if err != nil {
	// 	logger.Error("cannot parse root template",
	// 		zap.String("error", err.Error()),
	// 	)
	// }
	// err = t.Execute(w, r)
	// if err != nil {
	// 	logger.Error("cannot execute root template",
	// 		zap.String("error", err.Error()),
	// 	)
	// }

	// threads, err := data.Threads()
	// if err != nil {
	// error_message(writer, request, "Cannot get threads")
	// } else {
	// _, err := session(writer, request)
	// if err != nil {
	generateHTML(w, nil, "layout", "navbar-public", "index")
	// } else {
	// generateHTML(writer, threads, "layout", "private.navbar", "index")
	// }
	// }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout-no-header", "login")
}

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["email"][0]
	password := r.Form["password"][0]

	user, err := db.UserByEmail(email)
	if err != nil {
		log.Error(err.Error())
		return
	}

	if db.CheckPasswordHash(password, user.Password) {
		_, err := user.CreateSession()
		if err != nil {
			fmt.Fprintln(w, err.Error())
			log.Error(err.Error())
			return
		}
		r.Response.Re
	} else {
		// Redirect to login page
	}
}
