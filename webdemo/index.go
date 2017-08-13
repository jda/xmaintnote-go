package main

import (
	"net/http"

	"github.com/golang/glog"
)

// show index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("request: %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)

	renderTemplate(w, "page_index.tmpl", nil)
}
