package xmaintnote

import (
  "os"
  "testing"
)

// Test parsing the maint note standard example ical
// TODO more test cases & switch to generic func with
// test fixtures
func TestParseMaintNoteExample(t * testing.T) {
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
