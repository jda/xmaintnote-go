// Package icalutils provides helpr libraries for working with Calendar structs
// like those provided by github.com/luxifer/ical
package icalutils

import (
	"github.com/luxifer/ical"
)

func GetPropVal(p []*ical.Property, name string) string {
	for _, prop := range p {
		if name == prop.Name {
			return prop.Value
		}
	}
	return ""
}
