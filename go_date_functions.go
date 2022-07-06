package go_date_functions

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

//AddHours Add the specified number of hours to the given date.
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
	for i := 1; i <= amount; i++ {
		if AddDays(date, 1).Weekday() == time.Saturday {
			date = AddDays(date, 2)
		}
		if AddDays(date, 1).Weekday() == time.Sunday {
			date = AddDays(date, 1)
		}
		date = AddDays(date, 1)
	}
	return date
}

func subtractBusinessDays(dirtyDate time.Time, amount int) time.Time {
	date := AddDays(dirtyDate, 0)
	for i := 1; i <= amount; i++ {
		if AddDays(date, -1).Weekday() == time.Sunday {
			date = AddDays(date, -2)
		}
		if AddDays(date, -1).Weekday() == time.Saturday {
			date = AddDays(date, -1)
		}
		date = AddDays(date, -1)
	}
	return date
}

//AddMonths Add the specified number of months to the given date.
func AddMonths(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	zeroEdDate := time.Date(dirtyDate.Year(), dirtyDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	endingMonth := time.Date(dirtyDate.Year(), zeroEdDate.Month()+time.Month(amount), 1, 0, 0, 0, 0, time.UTC)
	daysEndingMonthDays := countDays(endingMonth)
	if dirtyDate.Day() > daysEndingMonthDays {
		return time.Date(endingMonth.Year(), endingMonth.Month(), daysEndingMonthDays, dirtyDate.Hour(), dirtyDate.Minute(), dirtyDate.Second(), dirtyDate.Nanosecond(), time.UTC)
	}
	return dirtyDate.AddDate(0, amount, 0)
}

func countDays(date time.Time) int {
	count := 0
	zeroEdDate := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	nextMonth := time.Date(date.Year(), zeroEdDate.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	for zeroEdDate.Before(nextMonth) {
		count = count + 1
		zeroEdDate = zeroEdDate.AddDate(0, 0, 1)
	}
	return count
}
