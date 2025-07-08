package datefns

import "time"

// AddHours Add the specified number of hours to the given date.
func AddHours(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Hour * time.Duration(amount))
}
