package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/golang/glog"
)

var genHTML = `<html><head>
<title>Make a Maint notice</title>
</head><body>
<h1>Make a xmaintnote ical</h1><p>
<form method="POST">
Summary:
Provider:
Account:
Maintenance ID:
Sequence:
Organizer Email
Impact
Status
Start
End
Maintenance Objects (at least one)
Name
Data
<input type="submit" value="generate">
</form>
</p></body></html>`
var genTmpl = template.Must(template.New("gen_in").Parse(genHTML))

// handle generation and download
func genHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("request: %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	switch r.Method {
	case "GET":
		fmt.Fprint(w, genHTML)
	case "POST":
		uploadHandler(w, r)
	default:
		nothingHandler(w, r)
	}
}
