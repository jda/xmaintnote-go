package main

import (
	"flag"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/golang/glog"
)

var tmplDir = flag.String("tmpldir", "tmpl", "path to template folder")
var tmplDebug = flag.Bool("tmpldebug", false, "reload templates on every request")
var templates map[string]*template.Template
var tmplMutex = &sync.Mutex{}

func main() {
	addr := flag.String("http", ":8080", "listen on this intf/port")
	flag.Parse()

	tmpl, err := loadTemplates(*tmplDir)
	if err != nil {
		glog.Fatalf("could not load templates: %s", err)
	}
	templates = tmpl

	http.Handle("/", regenTmpl(http.HandlerFunc(indexHandler)))
	http.Handle("/parse", regenTmpl(http.HandlerFunc(parseHandler)))
	http.Handle("/gen", regenTmpl(http.HandlerFunc(genHandler)))
	http.Handle("/_ah/health", regenTmpl(http.HandlerFunc(healthCheckHandler)))
	http.Handle("/favicon.ico", regenTmpl(http.HandlerFunc(nothingHandler)))

	glog.Infof("Listening for requests on %s", *addr)
	glog.Fatal(http.ListenAndServe(*addr, nil))
}

// regen templates
func regenTmpl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *tmplDebug == true {
			tmpl, err := loadTemplates(*tmplDir)
			if err != nil {
				glog.Fatalf("could not load templates: %s", err)

			}
			tmplMutex.Lock()
			templates = tmpl
			tmplMutex.Unlock()
		}
		h.ServeHTTP(w, r)
	})
}

func loadTemplates(tmplDir string) (tmpl map[string]*template.Template, err error) {
	tmpl = make(map[string]*template.Template)

	// base template
	baseTemplate := filepath.Join(tmplDir, "base.tmpl")

	// load page templates
	pageTemplates, err := filepath.Glob(tmplDir + "/page_*.tmpl")
	if err != nil {
		return tmpl, err
	}

	for _, pageTmpl := range pageTemplates {
		glog.Infof("parsing page template: %s", pageTmpl)
		files := []string{pageTmpl, baseTemplate}
		parsedTmpl := template.Must(template.ParseFiles(files...))
		tmpl[filepath.Base(pageTmpl)] = parsedTmpl
	}

	if len(tmpl) < 1 {
		glog.Warningf("no templates parsed from %s", tmplDir)
	}

	return tmpl, nil
}

func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		glog.Errorf("missing template %s", name)
	}

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.ExecuteTemplate(w, "base.tmpl", data)
	if err != nil {
		glog.Errorf("template exec: %s", err)
	}
}
