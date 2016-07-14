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

* `(te *TimeEx) Midnight() TimeEx` - returns time at the beginning of the day - 12AM
* `(te *TimeEx) EndOfDay() TimeEx` - returns time at the end of the day - 11:59PM
* `(te *TimeEx) Noon() TimeEx` - returns time at the midday - 12PM
* `(te *TimeEx) StartOfWeek() TimeEx` - returns date at the beginning of the week
* `(te *TimeEx) EndOfWeek() TimeEx` - returns date at the end of the week
* `(te *TimeEx) StartOfMonth() TimeEx` - returns date at the start of the month
* `(te *TimeEx) EndOfMonth() TimeEx` - returns date at the end of the month

## Usage
* install package: `go install github.com/alexchee/go-time-ex`
* import package:
```
import "github.com/alexchee/go-time-ex"
```
* create a `TimeEx` for a `Time`:
  ```
  te := TimeEx{Time: time.Now()}
  ```
* Call any of the methods and it'll return a `Time`:
  ex. `te.Midnight()`
* If you want to change the starting weekday for a week (default is Sunday) for calls like `EndOfWeek()` or `StartOfWeek()`, set `beginningWeekday` to that Weekday:
  ex. Use Monday as start of week.
  ```
  import "time"
  te := TimeEx{Time: time.Now(), beginningWeekday: time.Monday}
  ```

## Tests
`go test`

## Benchmark
`go test -bench=.`

## Contributing
This is my experiment with Go, so it's probably riddled with problems. Please open up a PR and improve this!

## TODO:
-[X] implement/fix `StartOfWeek`
-[ ] figure out a cleaner way to write test cases
-[ ] More time methods
-[ ] optimize methods
