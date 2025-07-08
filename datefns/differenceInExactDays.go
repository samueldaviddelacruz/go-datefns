package datefns

import "time"

// DifferenceInExactDays returns the exact number of days between two dates,
// including fractional days based on the time difference.
//
// A positive value indicates that laterDate is after earlierDate,
// and a negative value indicates the opposite.
//
// For example:
//
//	DifferenceInExactDays("2024-07-06T12:00", "2024-07-05T12:00") => 1.0
//	DifferenceInExactDays("2024-07-06T06:00", "2024-07-05T12:00") => 0.75
func DifferenceInExactDays(laterDate, earlierDate time.Time) float64 {
	return laterDate.Sub(earlierDate).Hours() / 24
}
