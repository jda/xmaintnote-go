package xmaintnote

import (
	"os"
	"testing"
)

// Test parsing the maint note standard example ical
// TODO more test cases & switch to generic func with
// test fixtures
func TestParseMaintNoteExample(t *testing.T) {
	fname := "testdata/maint-note-std-example.ical"
	f, err := os.Open(fname)
	if err != nil {
		t.Error(err)
	}

	_, err = ParseMaintNote(f)
	if err != nil {
		t.Error(err)
	}
}
func TestValidImpact(t *testing.T) {
	vars := map[string]bool{
		"invalidstatus":         false,
		"123":                   false,
		ImpactNone:              true,
		ImpactReducedRedundancy: true,
		ImpactDegraded:          true,
		ImpactOutage:            true,
	}

	for impact, res := range vars {
		if ValidImpact(impact) != res {
			t.Errorf("impact '%s' failed test", impact)
		}
	}
}

func TestInvalidImpact(t *testing.T) {
	if ValidImpact("spam123") {
		t.Errorf("invalid impact allowed")
	}
}

func TestValidStatus(t *testing.T) {
	vars := map[string]bool{
		"invalid things": false,
		"439123":         false,
		StatusTenative:   true,
		StatusCancelled:  true,
		StatusInProcess:  true,
		StatusCompleted:  true,
	}

	for status, res := range vars {
		if ValidStatus(status) != res {
			t.Errorf("status '%s' failed test", status)
		}
	}
}

func TestInvalidStatus(t *testing.T) {
	if ValidStatus("examplebad") {
		t.Errorf("invalid status allowed")
	}
}
