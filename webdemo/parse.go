package main

import (
	"net/http"

	"github.com/golang/glog"
	xmaintnote "github.com/jda/xmaintnote-go"
)

func parseHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("request: %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	switch r.Method {
	case "GET":
		renderTemplate(w, "page_parse.tmpl", nil)
	case "POST":
		uploadHandler(w, r)
	default:
		nothingHandler(w, r)
	}
}

// handle uploads and show results
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	infile, header, err := r.FormFile("fname")
	if err != nil {
		http.Error(w, "Error fetching uploaded file: "+err.Error(), http.StatusBadRequest)
		glog.Infof("%s: %s", header.Filename, infile)
		return
	}

	tmplVars := make(map[string]interface{})
	mn, err := xmaintnote.ParseMaintNote(infile)
	if err != nil {
		tmplVars["Error"] = err.Error()
		glog.Errorf("could not parse uploaded file `%s`", header.Filename)
	}

	tmplVars["Maint"] = mn

	renderTemplate(w, "page_parse.tmpl", tmplVars)
}
