package datefns

import "time"

// AddMinutes Add the specified number of minutes to the given date
func AddMinutes(dirtyDate time.Time, amount int) time.Time {
	if amount == 0 {
		// If amount == 0, no-op
		return dirtyDate.Add(0)
	}
	return dirtyDate.Add(time.Minute * time.Duration(amount))
}
