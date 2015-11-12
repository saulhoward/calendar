# calendar

Go library for creating calendars, adding events to them, and comparing schedules.

## Example

```go
package main
import (
    "fmt"
    "time"
    
    "github.com/saulhoward/calendar"
)

func main() {
    bobsCal := calendar.New("Bob's holiday calendar")
    
    // create events
    eventStart := time.Date(2015, time.November, 5, 19, 0, 0, 0, time.UTC)
    event := calendar.Event{
        Start: eventStart,
        End:   eventStart.Add(2 * time.Hour),
        Id:    "Guy Fawkes Night",
    }
    bobsCal.AddEvent(event)

    // check schedules
    checkStart := time.Date(2015, time.November, 5, 20, 0, 0, 0, time.UTC)
    isAvailable := calendar.AvailableRange(
        checkStart,
        checkStart.Add(1 * time.Hour),
        bobsCal,
    )
    fmt.Printf("is Bob available = %v\n", isAvailable)
}
```

## Dev notes

iCalendar spec

http://tools.ietf.org/html/rfc5545
