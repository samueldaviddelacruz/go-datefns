package datefns

import (
	"math"
	"time"
)

type RoundingMethod int

const (
	RoundNone    RoundingMethod = iota // Return exact float (default)
	RoundDown                          // Truncate toward zero (floor)
	RoundUp                            // Always round up
	RoundNearest                       // Round to nearest integer
)

func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// DifferenceInDaysWithRounding returns the number of days between two dates using the given rounding method.
func DifferenceInDaysWithRounding(laterDate, earlierDate time.Time, method RoundingMethod) float64 {
	diff := DifferenceInExactDays(laterDate, earlierDate)
	switch method {
	case RoundDown:
		return math.Floor(diff)
	case RoundUp:
		return math.Ceil(diff)
	case RoundNearest:
		return math.Round(diff)
	default:
		return diff
	}
}
