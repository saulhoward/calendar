package calendar

import (
	"sort"
	"sync"
	"time"
)

type Event struct {
	Id    string
	Start time.Time
	End   time.Time
}

// byStartTime implements sort.Interface for []Event based on
// the Start field.
type byStartTime []Event

func (s byStartTime) Len() int           { return len(s) }
func (s byStartTime) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byStartTime) Less(i, j int) bool { return s[i].Start.Before(s[j].Start) }

type Calendar struct {
	events map[string]Event
	id     string
	mu     sync.Mutex
}

// New returns a new Calendar
func New(id string) *Calendar {
	cal := &Calendar{
		events: make(map[string]Event),
		id:     id,
	}
	return cal
}

// AddEvent adds an event to the calendar
func (c Calendar) AddEvent(e Event) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.events[e.Id] = e
}

// Events returns all events, ordered by start time
func (c Calendar) Events() []Event {
	evs := make([]Event, len(c.events))
	i := 0
	for _, ev := range c.events {
		evs[i] = ev
		i = i + 1
	}
	sort.Sort(byStartTime(evs))
	return evs
}

// Available checks calendars for events during range, and returns false
// if any of them have conflicting events
func AvailableRange(start, end time.Time, cals ...*Calendar) bool {
	for _, c := range cals {
		for _, e := range c.events {
			if rangeConflict(e.Start, e.End, start, end) {
				return false
			}
		}
	}
	return true
}

func rangeConflict(aStart, aEnd, bStart, bEnd time.Time) bool {
	return (bStart.After(aStart) && bStart.Before(aEnd)) || (bEnd.After(aStart) && bEnd.Before(aEnd))
}
