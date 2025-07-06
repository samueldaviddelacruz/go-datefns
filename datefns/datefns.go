package datefns

import (
	"time"
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
