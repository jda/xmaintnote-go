// maintcat is a CLI tool for viewing and creating xmaintnote ical files
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/jda/xmaintnote-go"
)

const eventTemplate = `
Maintenance Event
Summary: {{.Summary}}
Start: {{.Start}}
End: {{.End}}
Provider: {{.Provider}}
Account: {{.Account}}
Maintenance: {{.MaintenanceID}}
Impact: {{.Impact}}
Status: {{.Status}}
Organizer: {{.OrganizerEmail}}

Affected objects:
{{range .Objects}}
Name: {{.Name}}
Data: {{.Data}}
{{end}}
`

func usage() {
	fmt.Printf("%s: [options] [FILENAME]\n", os.Args[0])
	fmt.Println("Maintenance notice read from STDIN if FILENAME is not provided.")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// no args, stdin, arg, assume file and read it
	if len(flag.Args()) == 0 {
		r := bufio.NewReader(os.Stdin)
		dumpCal(r)
	} else {
		f, err := os.Open(flag.Args()[0])
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r := bufio.NewReader(f)

		dumpCal(r)
	}
}

func dumpCal(r *bufio.Reader) {
	t := template.Must(template.New("event").Parse(eventTemplate))

	for {
		mn, err := xmaintnote.ParseMaintNote(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		for i := range mn.Events {
			spew.Dump(mn.Events[i])
			err := t.Execute(os.Stdout, mn.Events[i])
			if err != nil {
				panic(err)
			}
		}
	}
}
