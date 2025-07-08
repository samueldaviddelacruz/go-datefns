package datefns

import "time"

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
