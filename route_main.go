package main

import (
	"fmt"
	"net/http"

	"github.com/afeeblechild/fulcrum/lib/db"
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
		fmt.Fprintln(w, err.Error())
		return
	}

	if db.CheckPasswordHash(password, user.Password) {
		fmt.Fprintln(w, "valid")
		// TODO invalid password prompt
		// Would be nice to not actually redirect to this handler until the password is confirmed
		// That might be too much right now
	} else {
		fmt.Fprintln(w, "not valid")
		// TODO add session to db
	}
}
