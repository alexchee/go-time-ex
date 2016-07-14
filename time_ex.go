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
func (te *TimeEx) Midnight() *TimeEx {
  t := te.Time
  te.Time = time.Date(t.Year(),t.Month(),t.Day(), 0, 0, 0, 0, t.Location())
  return te
}

func (te *TimeEx) BeginningOfDay() *TimeEx {
  return te.Midnight()
}

// Returns end of day
func (te *TimeEx) EndOfDay() *TimeEx {
  t := te.Time
  // return time.Date(t.Year(),t.Month(),t.Day(), 23, 59, 0, 0, t.Location())
  te.Time = time.Date(t.Year(),t.Month(),t.Day(), 23, 59, 0, 0, t.Location())
  return te
}

// Returns midday
func (te *TimeEx) Noon() *TimeEx {
  t := te.Time
  te.Time = time.Date(t.Year(),t.Month(),t.Day(), 12, 0, 0, 0, t.Location())
  return te
}
func (te *TimeEx) Midday() *TimeEx {
  return te.Noon()
}
func (te *TimeEx) MiddleOfDay() *TimeEx {
  return te.Noon()
}

// Return date of the start of the week
func (te *TimeEx) StartOfWeek() *TimeEx {
  t := te.Time
  daysOffset := ((int(t.Weekday()) - int(te.beginningWeekday)) + 7) % 7
  te.Time = t.AddDate(0, 0, -daysOffset)
  return te
}
func (te *TimeEx) BeginningOfWeek() *TimeEx {
  return te.StartOfWeek()
}

// Return date of last day of the week
func (te *TimeEx) EndOfWeek() *TimeEx {
  t := te.Time
  daysOffset := (6 - int(t.Weekday()) + int(te.beginningWeekday)) % 7
  te.Time = t.AddDate(0, 0, daysOffset)
  return te
}

// Return start date of the month
func (te *TimeEx) StartOfMonth() *TimeEx {
  t := te.Time
  // return time.Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
  te.Time = time.Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
  return te
}
func (te *TimeEx) BeginningOfMonth() *TimeEx {
  return te.StartOfMonth()
}

// Return last date of the month
func (te *TimeEx) EndOfMonth() *TimeEx {
  t := te.Time
  te.Time = time.Date(t.Year(), t.Month()+1, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location()).AddDate(0, 0, -1)
  return te
}

// Returns last weekday of the week, default Saturday
func (te *TimeEx) EndingWeekday() time.Weekday {
  if te.beginningWeekday == time.Sunday {
    return time.Saturday
  }

  return te.beginningWeekday - 1
}
