package xmaintnote

import (
	"fmt"
	"io"

	"github.com/jda/ical"
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

// ValidStatus checks if a status value is a valid X-MAINTNOTE status
func ValidStatus(status string) bool {
	if status == StatusTenative {
		return true
	}

	if status == StatusCancelled {
		return true
	}

	if status == StatusInProcess {
		return true
	}

	if status == StatusCompleted {
		return true
	}

	return false
}

// ValidImpact checks if a status value is a valid X-MAINTNOTE impact
func ValidImpact(impact string) bool {
	if impact == ImpactNone {
		return true
	}

	if impact == ImpactReducedRedundancy {
		return true
	}

	if impact == ImpactDegraded {
		return true
	}

	if impact == ImpactOutage {
		return true
	}

	return false
}
