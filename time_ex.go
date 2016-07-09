package timeEx

import (
  "time"
)

const dateFormat = "2016-02-08"
const dateTimeFormat = "2016-02-08 at 3:04pm"

type TimeEx struct {
  time.Time
  beginningWeekday time.Weekday // when to start beginningWeekday, defaults to Sunday
}

// Returns beginning of day
func (te *TimeEx) Midnight() time.Time {
  t := te.Time
  return time.Date(t.Year(),t.Month(),t.Day(), 0, 0, 0, 0, t.Location())
}

func (te *TimeEx) BeginningOfDay() time.Time {
  return te.Midnight()
}

// Returns end of day
func (te *TimeEx) EndOfDay() time.Time {
  t := te.Time
  return time.Date(t.Year(),t.Month(),t.Day(), 23, 59, 0, 0, t.Location())
}

// Returns midday
func (te *TimeEx) Noon() time.Time {
  t := te.Time
  return time.Date(t.Year(),t.Month(),t.Day(), 12, 0, 0, 0, t.Location())
}
func (te *TimeEx) Midday() time.Time {
  return te.Noon()
}
func (te *TimeEx) MiddleOfDay() time.Time {
  return te.Noon()
}

// Return date of the start of the week
func (te *TimeEx) StartOfWeek() time.Time {
  t := te.Time
  daysOffset := ((int(t.Weekday()) - int(te.beginningWeekday)) + 7) % 7
  return t.AddDate(0, 0, -daysOffset)
}
func (te *TimeEx) BeginningOfWeek() time.Time {
  return te.StartOfWeek()
}

// Return date of last day of the week
func (te *TimeEx) EndOfWeek() time.Time {
  t := te.Time
  daysOffset := (6 - int(t.Weekday()) + int(te.beginningWeekday)) % 7
  return t.AddDate(0, 0, daysOffset)
}

// Returns last weekday of the week, default Saturday
func (te *TimeEx) EndingWeekday() time.Weekday {
  if te.beginningWeekday == time.Sunday {
    return time.Saturday
  }

  return te.beginningWeekday - 1
}
