package xmaintnote

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/luxifer/ical"
)

// ErrEmptyCalendar is the error code returned when there are no events in
// calendar
const ErrEmptyCalendar string = "no events found in calendar"

// ErrNotAMaintEvent is the error code for events that are not valid
// maintenence events
const ErrNotAMaintEvent string = "not a valid maintenence event"

// ErrNoMaintEvents is the error code for when no maintenence events were found
// in the calendar
const ErrNoMaintEvents string = "no maintenence events in calendar"

// MaintNote represents a maintenance notice containing one or more
// maintnence events
type MaintNote struct {
	CalProdID  string // iCal product ID
	CalVersion string // iCal version
	CalMethod  string // iCal method
	Events     []MaintEvent
}

// MaintEvent represents a single maintenence event
type MaintEvent struct {
	Summary        string
	Start          time.Time
	End            time.Time
	Created        time.Time
	UID            string
	Sequence       int
	Provider       string
	Account        string
	MaintenanceID  string
	Objects        []MaintObject
	Impact         string
	Status         string
	OrganizerEmail string
}

// MaintObject represents the item that is the subject of the maintenence event
type MaintObject struct {
	Name string // Name of maintenence object
	Data string // Alternate Representation (URI or other data) of object
}

// IsValid checks if a MaintEvent represents a valid MaintEvent
// e.g. Has all required properties & those properties have valid values
func (me *MaintEvent) IsValid() (valid bool, err error) {
	if me.Start.IsZero() {
		return false, fmt.Errorf("no start time")
	}
	if me.End.IsZero() {
		return false, fmt.Errorf("no end time")
	}
	if me.Created.IsZero() {
		return false, fmt.Errorf("no creation timestamp")
	}

	if me.UID == "" {
		return false, fmt.Errorf("no UID")
	}

	if me.Summary == "" {
		return false, fmt.Errorf("no summary")
	}

	if me.OrganizerEmail == "" {
		return false, fmt.Errorf("no organizer email")
	}

	if me.Provider == "" {
		return false, fmt.Errorf("no provider")
	}

	if me.Account == "" {
		return false, fmt.Errorf("no account")
	}

	if me.MaintenanceID == "" {
		return false, fmt.Errorf("no maintenence ID")
	}

	if !ValidImpact(me.Impact) {
		return false, fmt.Errorf("invalid impact")
	}

	if !ValidStatus(me.Status) {
		return false, fmt.Errorf("invalid status")
	}

	if len(me.Objects) < 1 {
		return false, fmt.Errorf("no maintenance objects")
	}

	return true, nil
}

// NewMaintNote creates a new MaintNote instance
func NewMaintNote() *MaintNote {
	mn := MaintNote{}
	return &mn
}

// ParseMaintNote parses Maintenence Notification events from a reader
func ParseMaintNote(r io.Reader) (mn MaintNote, err error) {
	calendar, err := ical.Parse(r)
	if err != nil {
		return mn, err
	}

	mn, err = ParseCalendar(*calendar)
	if err != nil {
		return mn, err
	}

	fmt.Printf("MaintNote: %+v\n", mn)
	return mn, err
}

// ParseCalendar creates a MaintNote from a ical.Calendar
func ParseCalendar(ic ical.Calendar) (mn MaintNote, err error) {
	mn = MaintNote{}

	// get required caMETHODlendar properties
	mn.CalProdID = ic.Prodid
	mn.CalVersion = ic.Version
	mn.CalMethod = ic.Method

	// process events
	numEvents := len(ic.Events)
	// need at least one event
	if numEvents < 1 {
		return mn, errors.New(ErrEmptyCalendar)
	}

	// process events
	noMaintEvent := true
	for i := 0; i < numEvents; i++ {
		event, err := ParseEvent(*ic.Events[i])
		if err != nil && err.Error() == ErrNotAMaintEvent { // skip non-maint event
			continue
		} else if err != nil { // surface all other errors
			return mn, err
		}
		// got this far, then valid maint event
		noMaintEvent = false
		mn.Events = append(mn.Events, event)
	}

	if noMaintEvent == true {
		return mn, errors.New(ErrNoMaintEvents)
	}

	return mn, err
}

// ParseEvent loads maintenence event from a ical.Event
func ParseEvent(ie ical.Event) (me MaintEvent, err error) {
	me = MaintEvent{
		Summary: ie.Summary,
		Start:   ie.StartDate,
		End:     ie.EndDate,
		Created: ie.Timestamp,
		UID:     ie.UID,
	}

	rawOrgEmail := getPropVal(ie.Properties, "ORGANIZER")
	if rawOrgEmail != "" {
		me.OrganizerEmail = strings.Replace(rawOrgEmail, "mailto:", "", 1)
	}

	rawSequence := getPropVal(ie.Properties, "SEQUENCE")
	if rawSequence != "" {
		me.Sequence, err = strconv.Atoi(rawSequence)
		if err != nil {
			return me, err
		}
	}

	me.Objects = getMaintObjects(ie.Properties)

	me.Provider = getPropVal(ie.Properties, maintProvider)
	me.Account = getPropVal(ie.Properties, maintAccount)
	me.MaintenanceID = getPropVal(ie.Properties, maintMaintID)
	me.Impact = getPropVal(ie.Properties, maintImpact)
	me.Status = getPropVal(ie.Properties, maintStatus)

	_, err = me.IsValid()
	return me, err
}

func getMaintObjects(p []*ical.Property) (mo []MaintObject) {
	mo = []MaintObject{}
	for _, prop := range p {
		if prop.Name == maintObjectID {
			m := MaintObject{
				Name: prop.Value,
			}

			if val, ok := prop.Params["ALTREP"]; ok {
				if len(val.Values) == 1 {
					m.Data = val.Values[0]
				}
			}

			mo = append(mo, m)
		}
	}
	return mo
}

func getPropVal(p []*ical.Property, name string) string {
	for _, prop := range p {
		if name == prop.Name {
			return prop.Value
		}
	}
	return ""
}
