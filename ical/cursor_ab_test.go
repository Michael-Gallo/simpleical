package ical

import (
	"testing"

	"github.com/michael-gallo/simpleical/model"
)

// BenchmarkComponentResolve isolates the hot-path cost of resolving the current
// event/todo/alarm/timezone pointer many times (as on each property line).
func BenchmarkComponentResolve(b *testing.B) {
	cal := &model.Calendar{
		Events: []model.Event{{
			Alarms: []model.Alarm{{}},
		}},
		Todos: []model.Todo{{
			Alarms: []model.Alarm{{}},
		}},
		Journals:  []model.Journal{{}},
		FreeBusys: []model.FreeBusy{{}},
		TimeZones: []model.TimeZone{{
			Standard: []model.TimeZoneProperty{{}},
			Daylight: []model.TimeZoneProperty{{}},
		}},
	}
	cursor := componentCursor{
		event:    &cal.Events[0],
		todo:     &cal.Todos[0],
		journal:  &cal.Journals[0],
		freeBusy: &cal.FreeBusys[0],
		timeZone: &cal.TimeZones[0],
		alarm:    &cal.Events[0].Alarms[0],
		tzProp:   &cal.TimeZones[0].Standard[0],
	}

	b.Run("Index", func(b *testing.B) {
		var sink uintptr
		b.ReportAllocs()
		for b.Loop() {
			sink ^= resolveByIndex(cal, stateEvent)
			sink ^= resolveByIndex(cal, stateTodo)
			sink ^= resolveByIndex(cal, stateJournal)
			sink ^= resolveByIndex(cal, stateFreebusy)
			sink ^= resolveByIndex(cal, stateTimezone)
			sink ^= resolveByIndex(cal, stateEventAlarm)
			sink ^= resolveByIndex(cal, stateTodoAlarm)
			sink ^= resolveByIndex(cal, stateStandard)
		}
		_ = sink
	})

	b.Run("Cursor", func(b *testing.B) {
		var sink uintptr
		b.ReportAllocs()
		for b.Loop() {
			sink ^= resolveByCursor(&cursor, stateEvent)
			sink ^= resolveByCursor(&cursor, stateTodo)
			sink ^= resolveByCursor(&cursor, stateJournal)
			sink ^= resolveByCursor(&cursor, stateFreebusy)
			sink ^= resolveByCursor(&cursor, stateTimezone)
			sink ^= resolveByCursor(&cursor, stateEventAlarm)
			sink ^= resolveByCursor(&cursor, stateTodoAlarm)
			sink ^= resolveByCursor(&cursor, stateStandard)
		}
		_ = sink
	})
}

//go:noinline
func resolveByIndex(calendar *model.Calendar, currentState parserState) uintptr {
	switch currentState {
	case stateEventAlarm:
		return uintptr(len(calendar.Events[len(calendar.Events)-1].Alarms[len(calendar.Events[len(calendar.Events)-1].Alarms)-1].Description))
	case stateTodoAlarm:
		return uintptr(len(calendar.Todos[len(calendar.Todos)-1].Alarms[len(calendar.Todos[len(calendar.Todos)-1].Alarms)-1].Description))
	case stateEvent:
		return uintptr(len(calendar.Events[len(calendar.Events)-1].UID))
	case stateTimezone:
		return uintptr(len(calendar.TimeZones[len(calendar.TimeZones)-1].TimeZoneID))
	case stateTodo:
		return uintptr(len(calendar.Todos[len(calendar.Todos)-1].UID))
	case stateJournal:
		return uintptr(len(calendar.Journals[len(calendar.Journals)-1].UID))
	case stateFreebusy:
		return uintptr(len(calendar.FreeBusys[len(calendar.FreeBusys)-1].UID))
	case stateStandard:
		tz := &calendar.TimeZones[len(calendar.TimeZones)-1]
		return uintptr(len(tz.Standard[len(tz.Standard)-1].TimeZoneOffsetFrom))
	case stateDaylight:
		tz := &calendar.TimeZones[len(calendar.TimeZones)-1]
		return uintptr(len(tz.Daylight[len(tz.Daylight)-1].TimeZoneOffsetTo))
	case stateCalendar, stateFinished, stateSkipComponent:
		return 0
	default:
		return 0
	}
}

//go:noinline
func resolveByCursor(cursor *componentCursor, currentState parserState) uintptr {
	switch currentState {
	case stateEventAlarm, stateTodoAlarm:
		return uintptr(len(cursor.alarm.Description))
	case stateEvent:
		return uintptr(len(cursor.event.UID))
	case stateTimezone:
		return uintptr(len(cursor.timeZone.TimeZoneID))
	case stateTodo:
		return uintptr(len(cursor.todo.UID))
	case stateJournal:
		return uintptr(len(cursor.journal.UID))
	case stateFreebusy:
		return uintptr(len(cursor.freeBusy.UID))
	case stateStandard, stateDaylight:
		return uintptr(len(cursor.tzProp.TimeZoneOffsetFrom))
	case stateCalendar, stateFinished, stateSkipComponent:
		return 0
	default:
		return 0
	}
}

func TestCursorEscape(t *testing.T) {
	// Sanity: cursor fields are set from heap-backed calendar slices.
	cal := &model.Calendar{}
	cal.Events = append(cal.Events, model.Event{UID: "x"})
	var cursor componentCursor
	cursor.event = &cal.Events[len(cal.Events)-1]
	if cursor.event.UID != "x" {
		t.Fatal("cursor did not point at event")
	}
}
