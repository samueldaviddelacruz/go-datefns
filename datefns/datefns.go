package datefns

import (
	"math"
	"time"
)

type RoundingMethod int

const (
	RoundNone    RoundingMethod = iota // Return exact float (default)
	RoundDown                          // Truncate toward zero (floor)
	RoundUp                            // Always round up
	RoundNearest                       // Round to nearest integer
)

// AddMilliseconds Add the specified number of Milliseconds to the given date
func AddMilliseconds(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Millisecond * time.Duration(amount))
}

// AddMinutes Add the specified number of minutes to the given date
func AddMinutes(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Minute * time.Duration(amount))
}

// AddHours Add the specified number of hours to the given date.
func AddHours(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Hour * time.Duration(amount))
}

// AddDays Add the specified number of days to the given date.
func AddDays(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	return dirtyDate.AddDate(0, 0, amount)
}

// AddBusinessDays Add the specified number of business days (mon - fri) to the given date, ignoring weekends.
func AddBusinessDays(dirtyDate time.Time, amount int) time.Time {
	if amount < 0 {
		return subtractBusinessDays(dirtyDate, amount*-1)
	}
	date := AddDays(dirtyDate, 0)
	for i := 0; i < amount; {
		date = AddDays(date, 1)
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			i++
		}
	}
	return date
}

func subtractBusinessDays(dirtyDate time.Time, amount int) time.Time {
	date := AddDays(dirtyDate, 0)
	for i := 1; i <= amount; {
		date = AddDays(date, -1)
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			i++
		}
	}
	return date
}

func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// AddMonths Add the specified number of months to the given date.
func AddMonths(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	year, month := dirtyDate.Year(), dirtyDate.Month()
	targetMonth := month + time.Month(amount)
	targetDate := time.Date(year, targetMonth, 1, dirtyDate.Hour(), dirtyDate.Minute(), dirtyDate.Second(), dirtyDate.Nanosecond(), dirtyDate.Location())

	daysInTargetMonth := daysInMonth(targetDate.Year(), targetDate.Month())
	day := min(dirtyDate.Day(), daysInTargetMonth)

	return time.Date(targetDate.Year(), targetDate.Month(), day, dirtyDate.Hour(), dirtyDate.Minute(), dirtyDate.Second(), dirtyDate.Nanosecond(), dirtyDate.Location())
}

// IsWeekend Does the given date fall on a weekend? A weekend is either Saturday or Sunday
func IsWeekend(dirtyDate time.Time) bool {
	return dirtyDate.Weekday() == time.Saturday || dirtyDate.Weekday() == time.Sunday
}

// DifferenceInCalendarDays Get the number of full days between the given dates.
//
// A positive value indicates that laterDate is after earlierDate,
// and a negative value indicates the opposite.
func DifferenceInCalendarDays(laterDate time.Time, earlierDate time.Time) int {
	later := time.Date(laterDate.Year(), laterDate.Month(), laterDate.Day(), 0, 0, 0, 0, laterDate.Location())
	earlier := time.Date(earlierDate.Year(), earlierDate.Month(), earlierDate.Day(), 0, 0, 0, 0, earlierDate.Location())
	result := later.Sub(earlier).Hours() / 24
	return int(result)
}

// DifferenceInExactDays returns the exact number of days between two dates,
// including fractional days based on the time difference.
//
// A positive value indicates that laterDate is after earlierDate,
// and a negative value indicates the opposite.
//
// For example:
//
//	DifferenceInExactDays("2024-07-06T12:00", "2024-07-05T12:00") => 1.0
//	DifferenceInExactDays("2024-07-06T06:00", "2024-07-05T12:00") => 0.75
func DifferenceInExactDays(laterDate, earlierDate time.Time) float64 {
	return laterDate.Sub(earlierDate).Hours() / 24
}

// DifferenceInDaysWithRounding returns the number of days between two dates using the given rounding method.
func DifferenceInDaysWithRounding(laterDate, earlierDate time.Time, method RoundingMethod) float64 {
	diff := DifferenceInExactDays(laterDate, earlierDate)
	switch method {
	case RoundDown:
		return math.Floor(diff)
	case RoundUp:
		return math.Ceil(diff)
	case RoundNearest:
		return math.Round(diff)
	default:
		return diff
	}
}
