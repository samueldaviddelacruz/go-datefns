package datefns

import "time"

// AddDays Add the specified number of days to the given date.
func AddDays(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	return dirtyDate.AddDate(0, 0, amount)
}
