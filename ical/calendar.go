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

func validateCalendar(calendar *model.Calendar, cps *calendarParseState) error {
	if calendar.Version == "" {
		return icalerr.ErrMissingCalendarVersionProperty
	}
	if calendar.ProdID == "" {
		return icalerr.ErrMissingCalendarProdIDProperty
	}
	if !cps.sawComponent {
		return icalerr.ErrMissingCalendarComponent
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
		for _, dt := range []model.DateTime{e.DTStamp, e.Start, e.Created, e.LastModified, e.RecurrenceID, e.End} {
			if err := check(dt); err != nil {
				return err
			}
		}
		for _, dt := range e.ExceptionDates {
			if err := check(dt); err != nil {
				return err
			}
		}
		for _, rd := range e.Rdate {
			if rd.DateTime != nil {
				if err := check(*rd.DateTime); err != nil {
					return err
				}
			}
			if rd.Period != nil {
				if err := check(rd.Period.Start); err != nil {
					return err
				}
				if err := check(rd.Period.End); err != nil {
					return err
				}
			}
		}
	}
	for i := range calendar.Todos {
		t := &calendar.Todos[i]
		for _, dt := range []model.DateTime{t.DTStamp, t.Completed, t.Created, t.DTStart, t.Due, t.LastModified, t.RecurrenceID} {
			if err := check(dt); err != nil {
				return err
			}
		}
		for _, dt := range t.ExceptionDates {
			if err := check(dt); err != nil {
				return err
			}
		}
		for _, rd := range t.Rdate {
			if rd.DateTime != nil {
				if err := check(*rd.DateTime); err != nil {
					return err
				}
			}
		}
	}
	for i := range calendar.Journals {
		j := &calendar.Journals[i]
		for _, dt := range []model.DateTime{j.DTStamp, j.Created, j.DTStart, j.LastModified, j.RecurrenceID} {
			if err := check(dt); err != nil {
				return err
			}
		}
		for _, dt := range j.ExceptionDates {
			if err := check(dt); err != nil {
				return err
			}
		}
		for _, rd := range j.Rdate {
			if rd.DateTime != nil {
				if err := check(*rd.DateTime); err != nil {
					return err
				}
			}
		}
	}
	for i := range calendar.FreeBusys {
		fb := &calendar.FreeBusys[i]
		for _, dt := range []model.DateTime{fb.DTStamp, fb.DTStart, fb.DTEnd} {
			if err := check(dt); err != nil {
				return err
			}
		}
	}
	_ = check
	return nil
}
