package main

import (
	"net/http"
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
