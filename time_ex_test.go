package timeEx

import (
	"testing"
	"time"
	"math/rand"
)

func TestMidnight(t *testing.T) {
	for _, test := range generateMidnightCases() {
		testTime := TimeEx{Time: test.startTime}
		observed := testTime.Midnight().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).Midnight() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestBeginningOfDay(t *testing.T) {
	for _, test := range generateMidnightCases() {
		testTime := TimeEx{Time: test.startTime}
		observed := testTime.BeginningOfDay().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).BeginningOfDay() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestEndOfDay(t *testing.T) {
	for _, test := range generateEndOfDayCases() {
		testTime := TimeEx{Time: test.startTime}
		observed := testTime.EndOfDay().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).BeginningOfDay() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestNoon(t *testing.T) {
	for _, test := range generateNoonCases() {
		testTime := TimeEx{Time: test.startTime}
		observed := testTime.Noon().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).BeginningOfDay() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestEndOfWeekStartingSunday(t *testing.T) {
	for _, test := range generateEndOfWeekSundayCases() {
		testTime := TimeEx{Time: test.startTime}
		observed := testTime.EndOfWeek().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).EndOfWeek() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestEndOfWeekStartingMonday(t *testing.T) {
	for _, test := range generateEndOfWeekMondayCases() {
		testTime := TimeEx{Time: test.startTime, beginningWeekday: time.Monday}
		observed := testTime.EndOfWeek().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).EndOfWeek() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestEndOfWeekStartingSaturday(t *testing.T) {
	for _, test := range generateEndOfWeekSaturdayCases() {
		testTime := TimeEx{Time: test.startTime, beginningWeekday: time.Saturday}
		observed := testTime.EndOfWeek().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).EndOfWeek() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestStartOfWeekStartingSunday(t *testing.T) {
	for _, test := range generateStartOfWeekSundayCases() {
		testTime := TimeEx{Time: test.startTime}
		observed := testTime.StartOfWeek().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).StartOfWeek() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestBeginningOfWeekStartingSunday(t *testing.T) {
	for _, test := range generateStartOfWeekSundayCases() {
		testTime := TimeEx{Time: test.startTime}
		observed := testTime.BeginningOfWeek().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).BeginningOfWeek() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestStartOfWeekStartingMonday(t *testing.T) {
	for _, test := range generateStartOfWeekMondayCases() {
		testTime := TimeEx{Time: test.startTime, beginningWeekday: time.Monday}
		observed := testTime.StartOfWeek().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).StartOfWeek() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestStartOfWeekStartingWednesday(t *testing.T) {
	for _, test := range generateStartOfWeekWednesdayCases() {
		testTime := TimeEx{Time: test.startTime, beginningWeekday: time.Wednesday}
		observed := testTime.StartOfWeek().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).StartOfWeek() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestBeginningOfMonth(t *testing.T) {
	for _, test := range generateStartOfMonthCases() {
		testTime := TimeEx{Time: test.startTime, beginningWeekday: time.Monday}
		observed := testTime.BeginningOfMonth().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).BeginningOfMonth() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestStartOfMonth(t *testing.T) {
	for _, test := range generateStartOfMonthCases() {
		testTime := TimeEx{Time: test.startTime, beginningWeekday: time.Monday}
		observed := testTime.StartOfMonth().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).StartOfMonth() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestEndOfMonth(t *testing.T) {
	for _, test := range generateEndOfMonthCases() {
		testTime := TimeEx{Time: test.startTime, beginningWeekday: time.Monday}
		observed := testTime.EndOfMonth().Time
		if !observed.Equal(test.expectedTime) {
			t.Fatalf("(%v).EndOfMonth() = %v, want %v (%s)",
				test.startTime, observed, test.expectedTime, test.description)
		}
	}
}

func TestEndingWeekday(t *testing.T) {
	for _, test := range testEndingWeekdayCases {
		testTime := TimeEx{beginningWeekday: test.beginningWeekday}
		observed := testTime.EndingWeekday()
		if observed != test.expectedWeekday {
			t.Fatalf("(%v).EndOfWeek() = %v, want %v (%s)",
				test.beginningWeekday, observed, test.expectedWeekday, test.description)
		}
	}
}

// Benchmark Midnight()
func BenchmarkMidnight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := TimeEx{Time: time.Now()}
		t.Midnight()
	}
}

// Benchmark Noon()
func BenchmarkNoon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := TimeEx{Time: time.Now()}
		t.Noon()
	}
}

// Benchmark EndOfDay()
func BenchmarkEndOfDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := TimeEx{Time: time.Now()}
		t.EndOfDay()
	}
}

// Benchmark EndOfWeek()
func BenchmarkEndOfWeek(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		t := TimeEx{Time: time.Now(), beginningWeekday: time.Weekday(rand.Intn(7))}
		t.EndOfWeek()
	}
}
