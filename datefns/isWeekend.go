package datefns

import "time"

// IsWeekend Does the given date fall on a weekend? A weekend is either Saturday or Sunday
func IsWeekend(dirtyDate time.Time) bool {
	return dirtyDate.Weekday() == time.Saturday || dirtyDate.Weekday() == time.Sunday
}
