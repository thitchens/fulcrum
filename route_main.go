package main

import (
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate := "web/templates/index.html"
	t, err := template.ParseFiles(rootTemplate)
	if err != nil {
		logger.Error("cannot parse root template",
			zap.String("error", err.Error()),
		)
	}
	err = t.Execute(w, r)
	if err != nil {
		logger.Error("cannot execute root template",
			zap.String("error", err.Error()),
		)
	}
}