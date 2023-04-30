package main

import (
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
	// TODO Get username/password from html

	user, err := db.UserByEmail(email)
	if err != nil {
		// TODO print error to page
	}

	if db.Encrypt(password) != user.Password {
		// TODO invalid password prompt
		// Would be nice to not actually redirect to this handler until the password is confirmed
		// That might be too much right now
	}else {
		// TODO add session to db
	}
}