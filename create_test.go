package xmaintnote

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var MaintEventTestCase1 = MaintEvent{
	Summary:        "Test Maint Event",
	Start:          (time.Now().Add(time.Hour * 36).Truncate(time.Second)),
	End:            (time.Now().Add(time.Hour * 48).Truncate(time.Second)),
	Created:        time.Now().Truncate(time.Second),
	UID:            "31336",
	Sequence:       1,
	Provider:       "Acme Internet",
	Account:        "Bugs Bunny",
	MaintenanceID:  "YOLO#1",
	Impact:         ImpactReducedRedundancy,
	Status:         StatusTenative,
	OrganizerEmail: "efudd@example.net",
}

// Create new NetMaint event
func TestCreateNewEvent(t *testing.T) {
	mo := MaintObject{
		Name: "yolocircuit#1",
	}
	MaintEventTestCase1.Objects = append(MaintEventTestCase1.Objects, mo)
	if ok, err := MaintEventTestCase1.IsValid(); ok != true {
		t.Error(err)
	}
}

// Export generated event to ical, parse, and verify
// same data
func TestGenerateParseEqual(t *testing.T) {
	if ok, err := MaintEventTestCase1.IsValid(); ok != true {
		t.Error(err)
	}
	mn := NewMaintNote()
	mn.Events = append(mn.Events, MaintEventTestCase1)

	data := mn.Export()

	buf := bytes.NewReader(data)
	newMN, err := ParseMaintNote(buf)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(&newMN, mn); diff != "" {
		t.Fatalf("generated maintnote ical does not match origional ical: (-got +want)\n%s", diff)
	}
}
