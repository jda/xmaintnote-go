package icalutils

import (
	"testing"

	"github.com/luxifer/ical"
)

var propVals = []*ical.Property{
	&ical.Property{
		Name:  "hello",
		Value: "world",
	},
}

func TestGetPropValOK(t *testing.T) {
	want := "world"
	pval := GetPropVal(propVals, "hello")
	if pval != want {
		t.Errorf("Expected val `%s`, got `%s`", want, pval)
	}
}

func TestGetPropValNo(t *testing.T) {
	want := ""
	pval := GetPropVal(propVals, "spamalot")
	if pval != want {
		t.Errorf("Expected val `%s`, got `%s`", want, pval)
	}
}
