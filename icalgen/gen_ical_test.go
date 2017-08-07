package icalgen

import (
	"bytes"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/luxifer/ical"
)

var testCaseNoXMaint = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Maint Note//https://github.com/maint-notification//
BEGIN:VEVENT
SUMMARY:Maint Note Example
DTSTART;VALUE=DATE-TIME:20160612T200000
DTEND;VALUE=DATE-TIME:20160612T210000
DTSTAMP;VALUE=DATE-TIME:20160612T192352Z
UID:42
SEQUENCE:1
ORGANIZER;CN="Example NOC":mailto:noone@example.com
END:VEVENT
END:VCALENDAR`

var testCaseWithXMaint = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Maint Note//https://github.com/maint-notification//
BEGIN:VEVENT
SUMMARY:Maint Note Example
DTSTART;VALUE=DATE-TIME:20160612T200000
DTEND;VALUE=DATE-TIME:20160612T210000
DTSTAMP;VALUE=DATE-TIME:20160612T192352Z
UID:42
SEQUENCE:1
X-MAINTNOTE-PROVIDER:example.com
X-MAINTNOTE-ACCOUNT:137.035999173
X-MAINTNOTE-MAINTENANCE-ID:WorkOrder-31415
X-MAINTNOTE-OBJECT-ID;ALTREP="https://example.org/maintenance?id=acme-widg
 ets-as-a-service":acme-widgets-as-a-service
X-MAINTNOTE-IMPACT:NO-IMPACT
X-MAINTNOTE-STATUS:TENTATIVE
ORGANIZER;CN="Example NOC":mailto:noone@example.com
END:VEVENT
END:VCALENDAR
`

// test generating and then parsing a ical without xmaintnote extensions
func TestGenerateICal(t *testing.T) {
	doTestCase(t, testCaseNoXMaint)
}

// test generating and then parsing a ical with xmaintnote extensions
func TestGenerateXMaintNoteICal(t *testing.T) {
	doTestCase(t, testCaseWithXMaint)
}

func doTestCase(t *testing.T, testCase string) {
	// test setup
	lfTestCase := strings.Replace(testCase, "\n", "\r\n", -1)
	inIcalReader := strings.NewReader(lfTestCase)
	inIcal, err := ical.Parse(inIcalReader)
	if err != nil {
		t.Fatalf("ical parse error on test setup: %s", err)
	}

	// test output
	outIcal := Export(inIcal)

	// parse results from output and compare to origional parse
	generatedIcalReader := bytes.NewReader(outIcal)
	generatedIcal, err := ical.Parse(generatedIcalReader)
	if err != nil {
		t.Fatalf("could not parse generated ical: %s", err)
	}
	spew.Dump(generatedIcal)

	if diff := cmp.Diff(inIcal, generatedIcal); diff != "" {
		t.Fatalf("generated ical does not match origional ical: (-got +want)\n%s", diff)
	}
}
