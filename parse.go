package xmaintnote

import (
	"fmt"
	"io"

	"github.com/luxifer/ical"
)

type MaintNote struct {
}

type MaintEvent struct {
	Summary       string
	Start         string
	End           string
	Created       string
	UID           string
	Sequence      int
	Provider      string
	Account       string
	MaintenanceID string
	ObjectID      string // CircuitID
	Impact        string
	Status        string
}

// NewMaintNote creates a new MaintNote instance
func NewMaintNote() *MaintNote {
	mn := MaintNote{}
	return &mn
}

// ParseMaintNote parses a Maintenence Notification from a reader
func ParseMaintNote(r io.Reader) (mn *MaintNote, err error) {
	err = parseIcal(r)
	return nil, err
}

func parseIcal(r io.Reader) error {
	calendar, err := ical.Parse(r)
	if err != nil {
		return err
	}

	fmt.Printf("Cal: %+v\n", calendar)
	printProp(calendar.Properties)
	printEvent(calendar.Events)
	return nil
}

func printProp(props []*ical.Property) {
	for p := range props {
		fmt.Printf("\tGot prop: %+v\n", props[p])
	}
}

func printEvent(events []*ical.Event) {
	for e := range events {
		fmt.Println("events!")
		printProp(events[e].Properties)
	}
}
