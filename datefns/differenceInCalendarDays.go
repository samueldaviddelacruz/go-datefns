package datefns

import "time"

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
