package calendar_test

import (
	"testing"
	"time"

	"github.com/saulhoward/calendar"
)

func TestAvailableRange(t *testing.T) {
	bobsCal := calendar.New("Bob's holiday calendar")

	eventStart := time.Date(2015, time.November, 5, 19, 0, 0, 0, time.UTC)
	event := calendar.Event{
		Start: eventStart,
		End:   eventStart.Add(2 * time.Hour),
		Id:    "Guy Fawkes Night",
	}
	bobsCal.AddEvent(event)

	conflictingEventStart := time.Date(2015, time.November, 5, 20, 0, 0, 0, time.UTC)
	isAvailable := calendar.AvailableRange(
		conflictingEventStart,
		conflictingEventStart.Add(1*time.Hour),
		bobsCal,
	)
	if isAvailable {
		t.Error("Conflicting range declared available")
	}

	nonconflictingEventStart := time.Date(2015, time.November, 6, 12, 0, 0, 0, time.UTC)
	isAvailable = calendar.AvailableRange(
		nonconflictingEventStart,
		nonconflictingEventStart.Add(1*time.Hour),
		bobsCal,
	)
	if !isAvailable {
		t.Error("Non conflicting range declared unavailable")
	}
}

func TestGetEvents(t *testing.T) {
	bobsCal := calendar.New("Bob's holiday calendar")

	eventStart := time.Date(2015, time.November, 5, 19, 0, 0, 0, time.UTC)
	event := calendar.Event{
		Start: eventStart,
		End:   eventStart.Add(2 * time.Hour),
		Id:    "later",
	}
	bobsCal.AddEvent(event)

	event2Start := time.Date(2015, time.November, 4, 19, 0, 0, 0, time.UTC)
	event2 := calendar.Event{
		Start: event2Start,
		End:   event2Start.Add(2 * time.Hour),
		Id:    "earlier",
	}
	bobsCal.AddEvent(event2)

	events := bobsCal.Events()

	if len(events) != 2 {
		t.Errorf("Wrong number of events returned. Expected 2, got %v", len(events))
	}

	if events[0].Id != "earlier" {
		t.Errorf("Events returned in the wrong order. Got %v first", events[0])
	}
}
