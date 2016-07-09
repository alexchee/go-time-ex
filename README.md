# go-time-ex
Go extensions for package `Time`. Inspired by Rails ActiveSupport Time extensions.

## `TimeEx` type
Adds `TimeEx` struct that adds some helper methods for `time`.
```
type TimeEx struct {
  time.Time
  beginningWeekday time.Weekday
}
```

## Helper methods

* `(te *TimeEx) Midnight() time.Time` - returns time at the beginning of the day - 12AM
* `(te *TimeEx) EndOfDay() time.Time` - returns time at the end of the day - 11:59PM
* `(te *TimeEx) Noon() time.Time` - returns time at the midday - 12PM
* `(te *TimeEx) StartOfWeek() time.Time` - returns date at the beginning of the week
* `(te *TimeEx) EndOfWeek() time.Time` - returns date at the end of the week

To change the weekday for beginning/end of week, set `TimeEx.beginningWeekday` to the weekday.

Ex. Setting it to be Monday:
```
currentTime := TimeEx{time: time.now(), beginningWeekday: time.Monday}
```

## Contributing
This is my experiment with Go, so it's probably riddled with problems. Please open up a PR and improve this!
