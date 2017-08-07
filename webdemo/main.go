package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/golang/glog"
	"github.com/jda/xmaintnote-go"
)

var parseHTML = `<html><head><title>xmaintnote parse</title></head><body>
<h1>xmaintnote parse results</h1><p>
{{range .Events}}
<h2>Event</h2>
<table border="1">
<tr><th>Summary</th><td>{{.Summary}}</td></tr>
<tr><th>Provider</th><td>{{.Provider}}</td></tr>
<tr><th>Account</th><td>{{.Account}}</td></tr>
<tr><th>Maintenance ID</th><td>{{.MaintenanceID}}</td></tr>
<tr><th>Sequence</th><td>{{.Sequence}}</td></tr>
<tr><th>Organizer Email</th><td>{{.OrganizerEmail}}</td></tr>
<tr><th>Impact</th><td>{{.Impact}}</td></tr>
<tr><th>Status</th><td>{{.Status}}</td></tr>
<tr><th>Created</th><td>{{.Created}}</td></tr>
<tr><th>Start</th><td>{{.Start}}</td></tr>
<tr><th>End</th><td>{{.End}}</td></tr>
<tr><td colspan="2">
Maintenance Objects
{{range .Objects}}
<table border="1" width="100%">
<tr><th>Name</th><td>{{.Name}}</td></tr>
<tr><th>Data</th><td>{{.Data}}</td></tr>
</table>
{{end}}
</td></tr>
</table>
{{end}}
</p></body></html>`
var parseTmpl = template.Must(template.New("parse_out").Parse(parseHTML))

func handler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("request: %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	switch r.Method {
	case "GET":
		indexHandler(w, r)
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

	mn, err := xmaintnote.ParseMaintNote(infile)
	if err != nil {
		http.Error(w, "Error parsing uploaded file: "+err.Error(), http.StatusBadRequest)
		glog.Errorf("could not parse uploaded file `%s`", header.Filename)
		return
	}

	err = parseTmpl.Execute(w, mn)
	if err != nil {
		glog.Errorf("err %s showing result for %s", err, header.Filename)
	}
}

// show main page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := `<html><head><title>xmaintnote-go demo</title></head><body>
<h1>xmaintnote-go demo</h1>
<p>Upload a xmaintnote-compatible ical file to continue.<br>
<form method="POST" enctype="multipart/form-data">
<input name="fname" type="file" />
<input type="submit" value="upload" />
</form></p>
</body></html>`

	fmt.Fprintf(w, page)
}

// does nothing, for sinking favicon or other common paths
// that we don't care about
func nothingHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	addr := flag.String("http", ":8080", "listen on this intf/port")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.HandleFunc("/favicon.ico", nothingHandler)

	glog.Infof("Listening for requests on %s", *addr)
	glog.Fatal(http.ListenAndServe(*addr, nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
