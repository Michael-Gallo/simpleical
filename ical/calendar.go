package ical

import (
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// parseCalendarProperty parses a single property line and sets its value in the provided vcalendar.
func parseCalendarProperty(propertyName string, value string, params map[string]string, calendar *model.Calendar) error {
	switch propertyName {
	case "VERSION":
		return setOnceProperty(&calendar.Version, value, propertyName, "VCALENDAR")
	case "PRODID":
		return setOnceProperty(&calendar.ProdID, value, propertyName, "VCALENDAR")
	case "CALSCALE":
		return setOnceProperty(&calendar.CalScale, value, propertyName, "VCALENDAR")
	case "METHOD":
		return setOnceProperty(&calendar.Method, value, propertyName, "VCALENDAR")
	default:
		appendExtensionProperty(&calendar.XProp, &calendar.IANAProp, propertyName, value, params)
		return nil
	}
}

// validateCalendar checks required VCALENDAR properties and that at least one component was seen.
// Event validation runs here so a late METHOD (accepted as a compatibility tolerance) can
// still influence METHOD-dependent VEVENT rules such as DTSTART.
func validateCalendar(calendar *model.Calendar, sawComponent bool) error {
	if calendar.Version == "" {
		return icalerr.ErrMissingCalendarVersionProperty
	}
	if calendar.ProdID == "" {
		return icalerr.ErrMissingCalendarProdIDProperty
	}
	if !sawComponent {
		return icalerr.ErrMissingCalendarComponent
	}
	for i := range calendar.Events {
		if err := validateEvent(&calendar.Events[i], calendar.Method); err != nil {
			return err
		}
	}
	return nil
}

// walkDates invokes fn for each DateTime in dates.
func walkDates(dates []model.DateTime, fn func(model.DateTime) error) error {
	for _, dt := range dates {
		if err := fn(dt); err != nil {
			return err
		}
	}
	return nil
}

// walkRDates invokes fn for each DateTime embedded in recurrence dates
// (DATE-TIME values and PERIOD start/end).
func walkRDates(rdates []model.RecurrenceDate, fn func(model.DateTime) error) error {
	for _, rd := range rdates {
		if rd.DateTime != nil {
			if err := fn(*rd.DateTime); err != nil {
				return err
			}
		}
		if rd.Period != nil {
			if err := fn(rd.Period.Start); err != nil {
				return err
			}
			if err := fn(rd.Period.End); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateCalendarTZIDs ensures VTIMEZONE TZIDs are unique and every LocalTZ
// DateTime references a defined VTIMEZONE.
func validateCalendarTZIDs(calendar *model.Calendar) error {
	tzids := make(map[string]struct{}, len(calendar.TimeZones))
	for i := range calendar.TimeZones {
		id := calendar.TimeZones[i].TimeZoneID
		if _, exists := tzids[id]; exists {
			return icalerr.ErrDuplicateTimezoneTZID
		}
		tzids[id] = struct{}{}
	}

	check := func(dt model.DateTime) error {
		if dt.Form != model.DateTimeFormLocalTZ {
			return nil
		}
		if _, ok := tzids[dt.TZID]; !ok {
			return icalerr.ErrUnknownTZID
		}
		return nil
	}

	for i := range calendar.Events {
		e := &calendar.Events[i]
		if err := walkDates([]model.DateTime{e.DTStamp, e.Start, e.Created, e.LastModified, e.RecurrenceID, e.End}, check); err != nil {
			return err
		}
		if err := walkDates(e.ExceptionDates, check); err != nil {
			return err
		}
		if err := walkRDates(e.Rdate, check); err != nil {
			return err
		}
	}
	for i := range calendar.Todos {
		t := &calendar.Todos[i]
		if err := walkDates([]model.DateTime{t.DTStamp, t.Completed, t.Created, t.DTStart, t.Due, t.LastModified, t.RecurrenceID}, check); err != nil {
			return err
		}
		if err := walkDates(t.ExceptionDates, check); err != nil {
			return err
		}
		if err := walkRDates(t.Rdate, check); err != nil {
			return err
		}
	}
	for i := range calendar.Journals {
		j := &calendar.Journals[i]
		if err := walkDates([]model.DateTime{j.DTStamp, j.Created, j.DTStart, j.LastModified, j.RecurrenceID}, check); err != nil {
			return err
		}
		if err := walkDates(j.ExceptionDates, check); err != nil {
			return err
		}
		if err := walkRDates(j.Rdate, check); err != nil {
			return err
		}
	}
	for i := range calendar.FreeBusys {
		fb := &calendar.FreeBusys[i]
		if err := walkDates([]model.DateTime{fb.DTStamp, fb.DTStart, fb.DTEnd}, check); err != nil {
			return err
		}
	}
	return nil
}
