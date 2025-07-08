package datefns

import "time"

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
