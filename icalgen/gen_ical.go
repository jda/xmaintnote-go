// Package icalgen provides just enough ical support to generate valid
// ical files from the subset of ical used by xmaintnote-go
package icalgen

import (
	"fmt"
	"strings"

	icu "github.com/jda/xmaintnote-go/icalutils"
	"github.com/luxifer/ical"
)

// TODO: swap the out, append, out, sprintf for func that encapsulates
// all that junk and handles wrapping of lines that are too long
// TODO use reader/writer interface?
// TODO clean up, properly implement RFC, then upstream

// Export exports a ical.Calendar in ical format as a byte array
func Export(c *ical.Calendar) (out []byte) {
	out = append(out, "BEGIN:VCALENDAR\r\n"...)

	version := icu.GetPropVal(c.Properties, "VERSION")
	if version != "" {
		out = append(out, fmt.Sprintf("VERSION:%s\r\n", version)...)
	}

	prodid := icu.GetPropVal(c.Properties, "PRODID")
	if prodid != "" {
		out = append(out, fmt.Sprintf("PRODID:%s\r\n", prodid)...)
	}

	for ei := range c.Events {
		event := c.Events[ei]

		out = append(out, "BEGIN:VEVENT\r\n"...)

		if event.Summary != "" {
			out = append(out, fmt.Sprintf("SUMMARY:%s\r\n", event.Summary)...)
		}

		if event.StartDate.IsZero() == false {
			d := event.StartDate.Format("20060102T150405")
			out = append(out, fmt.Sprintf("DTSTART;VALUE=DATE-TIME:%s\r\n", d)...)
		}

		if event.EndDate.IsZero() == false {
			d := event.EndDate.Format("20060102T150405")
			out = append(out, fmt.Sprintf("DTEND;VALUE=DATE-TIME:%s\r\n", d)...)
		}

		if event.Timestamp.IsZero() == false {
			d := event.Timestamp.UTC().Format("20060102T150405Z")
			out = append(out, fmt.Sprintf("DTSTAMP;VALUE=DATE-TIME:%s\r\n", d)...)
		}

		if event.UID != "" {
			out = append(out, fmt.Sprintf("UID:%s\r\n", event.UID)...)
		}

		sequence := icu.GetPropVal(event.Properties, "SEQUENCE")
		if sequence != "" {
			out = append(out, fmt.Sprintf("SEQUENCE:%s\r\n", sequence)...)
		}

		organizer := getFormatOrganizer(event)
		if organizer != "" {
			out = append(out, fmt.Sprintf("%s\r\n", organizer)...)
		}

		xProperties := genXProperties(event)
		out = append(out, xProperties...)

		out = append(out, "END:VEVENT\r\n"...)
	}

	out = append(out, "END:VCALENDAR\r\n"...)
	return out
}

func genXProperties(e *ical.Event) (seq []byte) {
	for pi := range e.Properties {
		prop := e.Properties[pi]
		if strings.HasPrefix(prop.Name, "X-") == true {
			line := prop.Name
			if len(prop.Params) == 0 {
				line += fmt.Sprintf(":%s", prop.Value)
			} else {
				for pmName := range prop.Params {
					param := prop.Params[pmName]
					line += fmt.Sprintf(";%s=\"%s\"", pmName, param.Values[0])
				}
				line += fmt.Sprintf(":%s", prop.Value)
			}
			seq = append(seq, fmt.Sprintf("%s\r\n", line)...)
		}
	}
	return seq
}

func getFormatOrganizer(e *ical.Event) (fo string) {
	for pi := range e.Properties {
		if e.Properties[pi].Name != "ORGANIZER" {
			continue
		}
		fo = "ORGANIZER"

		// TODO read spec and do the right thing here
		pr := e.Properties[pi]
		if v, ok := pr.Params["CN"]; ok {
			cn := v.Values[0]
			fo += fmt.Sprintf(`;CN="%s"`, cn)
		}
		email := pr.Value
		fo += fmt.Sprintf(":%s", email)
		return fo
	}
	return fo
}
