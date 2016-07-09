package timeEx

import (
  "time"
)

const testDateFormat = "2006-01-02"
const testDateTimeFormat = "2006-01-02 at 3:04pm"

type TimeExTestResult struct {
	startTime       time.Time
	expectedTime    time.Time
  description     string
}

type TimeExTestInput struct {
	startTime       string
	expectedTime    string
  description     string
}

type TimeExWeekdayTestResult struct {
  beginningWeekday time.Weekday
  expectedWeekday time.Weekday
  description     string
}

var midnightStrings = [...]TimeExTestInput{
	TimeExTestInput{"2016-02-08 at 10:04am", "2016-02-08 at 12:00am", "from morning"},
	TimeExTestInput{"2016-02-17 at 3:04pm", "2016-02-17 at 12:00am", "from afternoon"},
	TimeExTestInput{"2016-02-28 at 12:00pm", "2016-02-28 at 12:00am", "from noon"},
	TimeExTestInput{"2016-12-28 at 11:59pm", "2016-12-28 at 12:00am", "a minute before"},
	TimeExTestInput{"2016-12-28 at 12:01am", "2016-12-28 at 12:00am", "from minute after new day"},
}

var endOfDayStrings = [...]TimeExTestInput{
	TimeExTestInput{"2016-02-08 at 10:04am", "2016-02-08 at 11:59pm", "from morning"},
	TimeExTestInput{"2016-02-17 at 3:04pm", "2016-02-17 at 11:59pm", "from afternoon"},
	TimeExTestInput{"2016-02-28 at 12:00pm", "2016-02-28 at 11:59pm", "from noon"},
	TimeExTestInput{"2016-12-28 at 11:59pm", "2016-12-28 at 11:59pm", "exactly at end of day"},
	TimeExTestInput{"2016-12-28 at 12:01am", "2016-12-28 at 11:59pm", "from minute after new day"},
}

var noonStrings = [...]TimeExTestInput{
	TimeExTestInput{"2016-02-08 at 10:04am", "2016-02-08 at 12:00pm", "from morning"},
	TimeExTestInput{"2016-02-17 at 3:04pm", "2016-02-17 at 12:00pm", "from afternoon"},
	TimeExTestInput{"2016-02-28 at 12:00pm", "2016-02-28 at 12:00pm", "exactly at noon"},
	TimeExTestInput{"2016-12-28 at 11:59pm", "2016-12-28 at 12:00pm", "from end of day"},
	TimeExTestInput{"2016-12-28 at 12:01pm", "2016-12-28 at 12:00pm", "from minute after new noon"},
}

var endOfWeekSundayStrings = [...]TimeExTestInput{
  TimeExTestInput{"2016-02-07", "2016-02-13", "from beginning of week"},
  TimeExTestInput{"2016-02-17", "2016-02-20", "from middle of the week"},
  TimeExTestInput{"2016-02-27", "2016-02-27", "from end of the week"},
  TimeExTestInput{"2016-02-29", "2016-03-05", "goes to next month"},
  TimeExTestInput{"2015-12-30", "2016-01-02", "goes to next year"},
}

var endOfWeekMondayStrings = [...]TimeExTestInput{
  TimeExTestInput{"2016-02-08", "2016-02-14", "from beginning of week"},
  TimeExTestInput{"2016-02-17", "2016-02-21", "from middle of the week"},
  TimeExTestInput{"2016-02-28", "2016-02-28", "from end of the week"},
  TimeExTestInput{"2016-02-29", "2016-03-06", "goes to next month"},
  TimeExTestInput{"2016-12-28", "2017-01-01", "goes to next year"},
}

var endOfWeekSaturdayStrings = [...]TimeExTestInput{
  TimeExTestInput{"2016-02-06", "2016-02-12", "from beginning of week"},
  TimeExTestInput{"2016-02-17", "2016-02-19", "from middle of the week"},
  TimeExTestInput{"2016-02-26", "2016-02-26", "from end of the week"},
  TimeExTestInput{"2016-02-29", "2016-03-04", "goes to next month"},
  TimeExTestInput{"2015-12-30", "2016-01-01", "goes to next year"},
}

var startOfWeekMondayStrings = [...]TimeExTestInput{
  TimeExTestInput{"2016-02-08", "2016-02-08", "from beginning of week"},
  TimeExTestInput{"2016-02-10", "2016-02-08", "from middle of the week"},
  TimeExTestInput{"2016-02-14", "2016-02-08", "from end of the week"},
  TimeExTestInput{"2016-03-01", "2016-02-29", "goes to previous month"},
  TimeExTestInput{"2016-01-01", "2015-12-28", "goes to previous year"},
}

var testEndingWeekdayCases = []TimeExWeekdayTestResult{
  TimeExWeekdayTestResult{ time.Sunday, time.Saturday, "week begins on Sunday" },
  TimeExWeekdayTestResult{ time.Monday, time.Sunday, "week begins on Monday" },
  TimeExWeekdayTestResult{ time.Tuesday, time.Monday, "week begins on Tuesday" },
  TimeExWeekdayTestResult{ time.Wednesday, time.Tuesday, "week begins on Wednesday" },
  TimeExWeekdayTestResult{ time.Thursday, time.Wednesday, "week begins on Thursday" },
  TimeExWeekdayTestResult{ time.Friday, time.Thursday, "week begins on Friday" },
  TimeExWeekdayTestResult{ time.Saturday, time.Friday, "week begins on Saturday" },
}

// TODO figure out a neater way to do the Times from strings with one struct
func generateMidnightCases() []TimeExTestResult {
  testMidnightCases := []TimeExTestResult{}
  for _, midnightString := range midnightStrings {
    inputTime, _ := time.Parse(testDateTimeFormat, midnightString.startTime)
  	expectedTime, _ := time.Parse(testDateTimeFormat, midnightString.expectedTime)
  	testMidnightCases = append(testMidnightCases, TimeExTestResult{inputTime, expectedTime, midnightString.description})
	}
  return testMidnightCases
}

func generateEndOfDayCases() []TimeExTestResult {
  testEndOfDayCases := []TimeExTestResult{}
  for _, endOfDayString := range endOfDayStrings {
    inputTime, _ := time.Parse(testDateTimeFormat, endOfDayString.startTime)
  	expectedTime, _ := time.Parse(testDateTimeFormat, endOfDayString.expectedTime)
  	testEndOfDayCases = append(testEndOfDayCases, TimeExTestResult{inputTime, expectedTime, endOfDayString.description})
	}
  return testEndOfDayCases
}

func generateNoonCases() []TimeExTestResult {
  testNoonCases := []TimeExTestResult{}
  for _, noonStr := range noonStrings {
    inputTime, _ := time.Parse(testDateTimeFormat, noonStr.startTime)
  	expectedTime, _ := time.Parse(testDateTimeFormat, noonStr.expectedTime)
  	testNoonCases = append(testNoonCases, TimeExTestResult{inputTime, expectedTime, noonStr.description})
	}
  return testNoonCases
}


func generateBeginningOfWeekMondayCases() []TimeExTestResult {
  testBeginningOfWeekCases := []TimeExTestResult{}
  for _, startOfWeekStr := range startOfWeekMondayStrings {
    inputTime, _ := time.Parse(testDateFormat, startOfWeekStr.startTime)
  	expectedTime, _ := time.Parse(testDateFormat, startOfWeekStr.expectedTime)
  	testBeginningOfWeekCases = append(testBeginningOfWeekCases, TimeExTestResult{inputTime, expectedTime, startOfWeekStr.description})
	}

  return testBeginningOfWeekCases
}

func generateEndOfWeekMondayCases() []TimeExTestResult {
  testEndOfWeekCases := []TimeExTestResult{}
  for _, endOfWeekStr := range endOfWeekMondayStrings {
    inputTime, _ := time.Parse(testDateFormat, endOfWeekStr.startTime)
  	expectedTime, _ := time.Parse(testDateFormat, endOfWeekStr.expectedTime)
  	testEndOfWeekCases = append(testEndOfWeekCases, TimeExTestResult{inputTime, expectedTime, endOfWeekStr.description})
	}

  return testEndOfWeekCases
}

func generateEndOfWeekSundayCases() []TimeExTestResult {
  testEndOfWeekCases := []TimeExTestResult{}
  for _, endOfWeekStr := range endOfWeekSundayStrings {
    inputTime, _ := time.Parse(testDateFormat, endOfWeekStr.startTime)
  	expectedTime, _ := time.Parse(testDateFormat, endOfWeekStr.expectedTime)
  	testEndOfWeekCases = append(testEndOfWeekCases, TimeExTestResult{inputTime, expectedTime, endOfWeekStr.description})
	}

  return testEndOfWeekCases
}

func generateEndOfWeekSaturdayCases() []TimeExTestResult {
  testEndOfWeekCases := []TimeExTestResult{}
  for _, endOfWeekStr := range endOfWeekSaturdayStrings {
    inputTime, _ := time.Parse(testDateFormat, endOfWeekStr.startTime)
    expectedTime, _ := time.Parse(testDateFormat, endOfWeekStr.expectedTime)
    testEndOfWeekCases = append(testEndOfWeekCases, TimeExTestResult{inputTime, expectedTime, endOfWeekStr.description})
  }

  return testEndOfWeekCases
}
