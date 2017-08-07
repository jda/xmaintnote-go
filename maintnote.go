// Package xmaintnote parses and generates Maintenance BCOP-formatted
// iCalendar files to aid in coordinating network maintenance.
//
// Maintenance BCOP: https://www.facebook.com/groups/855738444449323/
package xmaintnote

import (
	"fmt"
	"reflect"
	"time"
)

//
// MaintNote fields
//
const maintProvider string = "X-MAINTNOTE-PROVIDER"
const maintAccount string = "X-MAINTNOTE-ACCOUNT"
const maintMaintID string = "X-MAINTNOTE-MAINTENANCE-ID"
const maintObjectID string = "X-MAINTNOTE-OBJECT-ID"
const maintImpact string = "X-MAINTNOTE-IMPACT"
const maintStatus string = "X-MAINTNOTE-STATUS"

//
// Impacts
//

// ImpactNone represents MAINTNOTE NO-IMPACT impact
const ImpactNone string = "NO-IMPACT"

// ImpactReducedRedundancy represents the MAINTNOTE REDUCED-REDUNDANCY impact
const ImpactReducedRedundancy string = "REDUCED-REDUNDANCY"

// ImpactDegraded represents the MAINTNOTE DEGRADED impact
const ImpactDegraded string = "DEGRADED"

// ImpactOutage represents the MAINTNOTE OUTAGE impact
const ImpactOutage string = "OUTAGE"

//
// Statuses
//

// StatusTenative represents the MAINTNOTE TENATIVE status
const StatusTenative string = "TENTATIVE"

// StatusCancelled represents the MAINTNOTE CANCELLED status
const StatusCancelled string = "CANCELLED"

// StatusInProcess represents the MAINTNOTE IN-PROCESS status
const StatusInProcess string = "IN-PROCESS"

// StatusCompleted represents the MAINTNOTE COMPLETED status
const StatusCompleted string = "COMPLETED"

// MaintNote represents a maintenance notice containing one or more
// maintnence events
type MaintNote struct {
	CalProdID  string // iCal product ID
	CalVersion string // iCal version
	CalMethod  string // iCal method
	Events     []MaintEvent
}

// NewMaintNote creates a new MaintNote instance
func NewMaintNote() *MaintNote {
	mn := MaintNote{
		CalVersion: "2.0",
	}

	mn.CalProdID = fmt.Sprintf("-//Maint Note//%s//", reflect.TypeOf(mn).PkgPath())

	return &mn
}

// MaintEvent represents a single maintenance event
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

// MaintObject represents the item that is the subject of the maintenance event
type MaintObject struct {
	Name string // Name of maintenance object
	Data string // Alternate Representation (URI or other data) of object
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
		return false, fmt.Errorf("no maintenance ID")
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
