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

// SubDays Substract the specified number of days to the given date.
func SubDays(dirtyDate time.Time, amount int) time.Time {
	return AddDays(dirtyDate, amount*-1)
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

// SubMonths Substract the specified number of months to the given date.
func SubMonths(dirtyDate time.Time, amount int) time.Time {
	return AddMonths(dirtyDate, amount*-1)
}

// AddYears Add the specified number of years to the given date.
func AddYears(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	aYearInMonths := amount * 12
	return AddMonths(dirtyDate, aYearInMonths)
}

// SubYears Substract the specified number of years to the given date.
func SubYears(dirtyDate time.Time, amount int) time.Time {
	return AddYears(dirtyDate, amount*-1)
}

// IsWeekend Does the given date fall on a weekend? A weekend is either Saturday or Sunday
func IsWeekend(dirtyDate time.Time) bool {
	return dirtyDate.Weekday() == time.Saturday || dirtyDate.Weekday() == time.Sunday
}

// IsToday Is the given date today?
func IsToday(dirtyDate time.Time, now time.Time) bool {
	tyear, tmonth, tday := now.Date()
	dyear, dmonth, dday := dirtyDate.Date()
	return tyear == dyear && tmonth == dmonth && tday == dday
}
func IsTodayNow(dirtyDate time.Time) bool {
	return IsToday(dirtyDate, time.Now())
}

// IsTomorrow Is the given date tomorrow?
func IsTomorrow(dirtyDate time.Time, now time.Time) bool {
	tyear, tmonth, tday := AddDays(now, 1).Date()
	dyear, dmonth, dday := dirtyDate.Date()
	return tyear == dyear && tmonth == dmonth && tday == dday
}

func IsTomorrowNow(dirtyDate time.Time) bool {
	return IsTomorrow(dirtyDate, time.Now())
}

// IsPast Is the given date in the past?
func IsPast(dirtyDate time.Time, now time.Time) bool {
	return dirtyDate.Before(now)
}
func IsPastNow(dirtyDate time.Time) bool {
	return dirtyDate.Before(time.Now())
}

// IsFuture Is the given date in the future?
func IsFuture(dirtyDate time.Time, now time.Time) bool {
	return now.Before(dirtyDate)
}
func IsFutureNow(dirtyDate time.Time) bool {
	return time.Now().Before(dirtyDate)
}

// IsSameDay Are the given dates in the same day (and year and month)?
func IsSameDay(laterDate time.Time, earlierDate time.Time) bool {
	return StartOfDay(laterDate).Equal(StartOfDay(earlierDate))
}

// IsSameMonth Are the given dates in the same month (and year)?
func IsSameMonth(laterDate time.Time, earlierDate time.Time) bool {
	return laterDate.Year() == earlierDate.Year() && laterDate.Month() == earlierDate.Month()
}

// IsSameYear Are the given dates in the same year?
func IsSameYear(laterDate time.Time, earlierDate time.Time) bool {
	return laterDate.Year() == earlierDate.Year()
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

// StartOfDay Return the start of a day for the given date.
func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// EndOfDay Return the end of a day for the given date.
func EndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999_999_999, date.Location())
}

// StartOfMonth start of a month for the given date.
func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

// EndOfMonth end of a month for the given date.
func EndOfMonth(date time.Time) time.Time {
	result := AddMonths(StartOfMonth(date), 1)
	result = AddDays(result, -1)
	return EndOfDay(result)
}
