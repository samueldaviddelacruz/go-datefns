package go_date_functions

import (
	"time"
)

// AddMilliseconds Add the specified number of Milliseconds to the given date
func AddMilliseconds(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If 0 days, no-op to avoid changing times in the hour before end of DST
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Millisecond * time.Duration(amount))
}

// AddMinutes Add the specified number of minutes to the given date
func AddMinutes(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If 0 days, no-op to avoid changing times in the hour before end of DST
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Minute * time.Duration(amount))
}

//AddHours Add the specified number of hours to the given date.
func AddHours(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If 0 days, no-op to avoid changing times in the hour before end of DST
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Hour * time.Duration(amount))
}

// AddDays Add the specified number of days to the given date.
func AddDays(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If 0 days, no-op to avoid changing times in the hour before end of DST
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add((time.Hour * 24) * time.Duration(amount))
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
