package main

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/google/uuid"
	xmaintnote "github.com/jda/xmaintnote-go"
)

const inputTimePattern string = "2006-01-02T15:04"

// handle generation and download
func genHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("request: %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	switch r.Method {
	case "GET":
		renderTemplate(w, "page_gen.tmpl", nil)
	case "POST":
		genFormHandler(w, r)
	default:
		nothingHandler(w, r)
	}
}

func genFormHandler(w http.ResponseWriter, r *http.Request) {
	tmplVars := make(map[string]interface{})

	mn := xmaintnote.NewMaintNote()
	maintEvent := xmaintnote.MaintEvent{
		Summary:        r.FormValue("summary"),
		Provider:       r.FormValue("provider"),
		Account:        r.FormValue("account"),
		MaintenanceID:  r.FormValue("maintid"),
		Impact:         r.FormValue("impact"),
		Status:         r.FormValue("status"),
		OrganizerEmail: r.FormValue("organizer"),
		Created:        time.Now().UTC(),
		UID:            uuid.New().String(),
	}

	startTime, err := time.Parse(inputTimePattern, r.FormValue("start"))
	if err != nil {
		glog.Warning("failed to parse startTime")
	}
	maintEvent.Start = startTime

	endTime, err := time.Parse(inputTimePattern, r.FormValue("end"))
	if err != nil {
		glog.Warning("failed to parse endTime")
	}
	maintEvent.End = endTime

	rawObjs := r.FormValue("maintobject")
	objs := strings.Split(rawObjs, "\r\n")
	for _, obj := range objs {
		parts := strings.SplitN(obj, ",", 2)
		if len(parts) != 2 {
			glog.Warning("ignoring maintobject, got wrong num of parts")
			continue
		}
		maintObj := xmaintnote.MaintObject{
			Name: parts[0],
			Data: parts[1],
		}
		maintEvent.Objects = append(maintEvent.Objects, maintObj)
	}

	_, err = maintEvent.IsValid()
	if err != nil {
		tmplVars["Error"] = err.Error()
		glog.Errorf("maintEvent is not valid: `%s`", err)
		renderTemplate(w, "page_gen.tmpl", tmplVars)
		return
	}

	mn.Events = []xmaintnote.MaintEvent{maintEvent}

	calBytes := mn.Export()
	cal := bytes.NewBuffer(calBytes)

	fname := maintEvent.MaintenanceID + ".ics"

	w.Header().Set("Content-Disposition", "attachment; filename="+fname)
	w.Header().Set("Content-Type", "text/calendar")

	cal.WriteTo(w)
}
