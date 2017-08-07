package xmaintnote

import (
	"strconv"

	"github.com/jda/xmaintnote-go/icalgen"
	"github.com/luxifer/ical"
)

// Export generates a ical file for the maintenance event
func (mn *MaintNote) Export() (cal []byte) {
	ic := ical.NewCalendar()

	calVersionProp := ical.Property{
		Name:  "VERSION",
		Value: mn.CalVersion,
	}
	ic.Properties = append(ic.Properties, &calVersionProp)

	calProdProp := ical.Property{
		Name:  "PRODID",
		Value: mn.CalProdID,
	}
	ic.Properties = append(ic.Properties, &calProdProp)

	for i := range mn.Events {
		calEvent := mn.Events[i].GetCalendarEvent()
		ic.Events = append(ic.Events, calEvent)
	}

	cal = icalgen.Export(ic)
	return cal
}

// GetCalendarEvent returns a ical Event representation of a Maintnence Event
func (me *MaintEvent) GetCalendarEvent() (ce *ical.Event) {
	ce = ical.NewEvent()

	ce.UID = me.UID
	ce.Summary = me.Summary

	ce.Timestamp = me.Created
	ce.StartDate = me.Start
	ce.EndDate = me.End

	ce.Properties = append(ce.Properties, &ical.Property{
		Name:  maintProvider,
		Value: me.Provider,
	})

	ce.Properties = append(ce.Properties, &ical.Property{
		Name:  maintAccount,
		Value: me.Account,
	})

	ce.Properties = append(ce.Properties, &ical.Property{
		Name:  maintMaintID,
		Value: me.MaintenanceID,
	})

	ce.Properties = append(ce.Properties, &ical.Property{
		Name:  maintImpact,
		Value: me.Impact,
	})

	ce.Properties = append(ce.Properties, &ical.Property{
		Name:  maintStatus,
		Value: me.Status,
	})

	ce.Properties = append(ce.Properties, &ical.Property{
		Name:  "ORGANIZER",
		Value: me.OrganizerEmail,
	})

	ce.Properties = append(ce.Properties, &ical.Property{
		Name:  "SEQUENCE",
		Value: strconv.Itoa(me.Sequence),
	})

	// not actually sure if spec allows more than one entry here, or if multiple
	// entries are only allowed at URL referenced by ALTREP...
	for i := range me.Objects {
		p := ical.Param{
			Values: []string{
				me.Objects[i].Data},
		}

		repr := ical.Property{
			Name:  maintObjectID,
			Value: me.Objects[i].Name,
			Params: map[string]*ical.Param{
				"ALTREP": &p,
			},
		}

		ce.Properties = append(ce.Properties, &repr)
	}

	return ce
}
