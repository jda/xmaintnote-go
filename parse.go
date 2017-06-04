package xmaintnote

import (
  "io"
  "fmt"
)

// MaintNote represents a Maintnance Notification
type MaintNote struct {

}

// NewMaintNote creates a new MaintNote instance
func NewMaintNote() *MaintNote {
  mn := MaintNote{}
  return &mn
}

// ParseMaintNote parses a Maintenence Notification from a reader
func ParseMaintNote(r io.Reader) (mn *MaintNote, err error) {
  err = fmt.Errorf("NOT IMPLEMENTED YET")
  return nil, nil
}
