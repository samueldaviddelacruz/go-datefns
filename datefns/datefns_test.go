package datefns

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestAddDays(t *testing.T) {

	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name      string
		dirtyDate time.Time
		amount    int
		want      time.Time
	}{
		{
			name:      "adds the given number of days",
			dirtyDate: testDate, amount: 1,
			want: time.Date(1991, 9, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "subtract the given number of days",
			dirtyDate: testDate, amount: -1,
			want: time.Date(1991, 9, 25, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDays(tt.dirtyDate, tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDays() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddDays(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddDays() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})

}

func TestAddBusinessDays(t *testing.T) {

	tests := []struct {
		name      string
		dirtyDate time.Time
		amount    int
		want      time.Time
	}{
		{
			name:      "adds the given number of business days",
			dirtyDate: time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC), amount: 10,
			want: time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "subtract the given number of business days",
			dirtyDate: time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC), amount: -1,
			want: time.Date(2022, 6, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "start on Saturday and add 1 business day",
			dirtyDate: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), amount: 1,
			want: time.Date(2023, 7, 3, 0, 0, 0, 0, time.UTC), // Monday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBusinessDays(tt.dirtyDate, tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBusinessDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddHours(t *testing.T) {

	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name      string
		dirtyDate time.Time
		amount    int
		want      time.Time
	}{
		{
			name:      "adds the given number of hours",
			dirtyDate: testDate, amount: 1,
			want: time.Date(1991, 9, 26, 1, 0, 0, 0, time.UTC),
		}, {
			name:      "subtract the given number of hours",
			dirtyDate: time.Date(1991, 9, 26, 1, 0, 0, 0, time.UTC), amount: -1,
			want: time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddHours(tt.dirtyDate, tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddHours() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddHours(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddHours() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})

}

func TestAddMinutes(t *testing.T) {

	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name      string
		dirtyDate time.Time
		amount    int
		want      time.Time
	}{
		{
			name:      "adds the given number of minutes",
			dirtyDate: testDate, amount: 1,
			want: time.Date(1991, 9, 26, 0, 1, 0, 0, time.UTC),
		},
		{
			name:      "subtract the given number of minutes",
			dirtyDate: time.Date(1991, 9, 26, 0, 1, 0, 0, time.UTC), amount: -1,
			want: time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMinutes(tt.dirtyDate, tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMinutes() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddMinutes(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddMinutes() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})
}

func TestAddMilliseconds(t *testing.T) {
	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name      string
		dirtyDate time.Time
		amount    int
		want      time.Time
	}{
		{
			name:      "adds the given number of milliseconds",
			dirtyDate: testDate, amount: 1000,
			want: time.Date(1991, 9, 26, 0, 0, 1, 0, time.UTC),
		}, {
			name:      "subtract the given number of milliseconds",
			dirtyDate: time.Date(1991, 9, 26, 0, 0, 1, 0, time.UTC), amount: -1000,
			want: time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMilliseconds(tt.dirtyDate, tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMilliseconds() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddMilliseconds(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddMilliseconds() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})
}

func TestAddMonths(t *testing.T) {

	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name      string
		dirtyDate time.Time
		amount    int
		want      time.Time
	}{
		{
			name:      "adds the given number of Months",
			dirtyDate: testDate, amount: 1,
			want: time.Date(1991, 10, 26, 0, 0, 0, 0, time.UTC),
		}, {
			name:      "subtract the given number of Months",
			dirtyDate: time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC), amount: -1,
			want: time.Date(1991, 8, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "adds the given number of Months (February)",
			dirtyDate: time.Date(2022, 1, 31, 0, 0, 0, 0, time.UTC), amount: 1,
			want: time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "adds month from Jan 31 to leap Feb",
			dirtyDate: time.Date(2020, 1, 31, 0, 0, 0, 0, time.UTC), amount: 1,
			want: time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMonths(tt.dirtyDate, tt.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMonths() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddMonths(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddMonths() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		name      string
		dirtyDate time.Time
		want      bool
	}{
		{
			"returns true if the given date is in a weekend",
			time.Date(2014, 10, 5, 0, 0, 0, 0, time.UTC),
			true,
		},
		{
			"returns false if the given date is not in a weekend",
			time.Date(2014, 10, 6, 0, 0, 0, 0, time.UTC),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWeekend(tt.dirtyDate); got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInCalendarDays(t *testing.T) {
	tests := []struct {
		name    string
		later   time.Time
		earlier time.Time
		want    int
	}{
		{
			name:    "returns the number of full days between the given dates",
			later:   time.Date(2012, 7, 2, 18, 0, 0, 0, time.UTC),
			earlier: time.Date(2011, 7, 2, 6, 0, 0, 0, time.UTC),
			want:    366,
		},
		{
			name:    "returns no full days between the given dates",
			later:   time.Date(2024, 7, 6, 5, 0, 0, 0, time.UTC),
			earlier: time.Date(2024, 7, 6, 4, 59, 0, 0, time.UTC),
			want:    0,
		},
		{
			name:    "returns the number of full days between the given dates in negative",
			later:   time.Date(2011, 7, 2, 6, 0, 0, 0, time.UTC),
			earlier: time.Date(2012, 7, 2, 18, 0, 0, 0, time.UTC),
			want:    -366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceInCalendarDays(tt.later, tt.earlier); got != tt.want {
				t.Errorf("DifferenceInDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInExactDays(t *testing.T) {
	tests := []struct {
		name    string
		later   time.Time
		earlier time.Time
		want    float64
	}{
		{
			name:    "exactly one day apart",
			later:   time.Date(2024, 7, 6, 12, 0, 0, 0, time.UTC),
			earlier: time.Date(2024, 7, 5, 12, 0, 0, 0, time.UTC),
			want:    1.0,
		},
		{
			name:    "half a day apart",
			later:   time.Date(2024, 7, 6, 12, 0, 0, 0, time.UTC),
			earlier: time.Date(2024, 7, 6, 0, 0, 0, 0, time.UTC),
			want:    0.5,
		},
		{
			name:    "same moment",
			later:   time.Date(2024, 7, 6, 12, 0, 0, 0, time.UTC),
			earlier: time.Date(2024, 7, 6, 12, 0, 0, 0, time.UTC),
			want:    0.0,
		},
		{
			name:    "negative result (laterDate is before earlierDate)",
			later:   time.Date(2024, 7, 5, 12, 0, 0, 0, time.UTC),
			earlier: time.Date(2024, 7, 6, 12, 0, 0, 0, time.UTC),
			want:    -1.0,
		},
		{
			name:    "one and a quarter days apart",
			later:   time.Date(2024, 7, 6, 18, 0, 0, 0, time.UTC),
			earlier: time.Date(2024, 7, 5, 12, 0, 0, 0, time.UTC),
			want:    1.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInExactDays(tt.later, tt.earlier)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("DifferenceInExactDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInDaysWithRounding(t *testing.T) {
	start := time.Date(2024, 7, 5, 12, 0, 0, 0, time.UTC)
	end := time.Date(2024, 7, 6, 18, 0, 0, 0, time.UTC) // 1.25 days apart

	tests := []struct {
		name   string
		method RoundingMethod
		want   float64
	}{
		{"RoundNone", RoundNone, 1.25},
		{"RoundDown", RoundDown, 1.0},
		{"RoundUp", RoundUp, 2.0},
		{"RoundNearest", RoundNearest, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInDaysWithRounding(end, start, tt.method)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Method %v: got %v, want %v", tt.method, got, tt.want)
			}
		})
	}
}

func TestStartOfDay(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want time.Time
	}{
		{
			"returns the date with the time set to 00:00:00",
			time.Date(2024, 7, 5, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 7, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			"keeps the original time zone",
			time.Date(2024, 7, 5, 23, 59, 59, 123456789, time.FixedZone("EST", -5*3600)),
			time.Date(2024, 7, 5, 0, 0, 0, 0, time.FixedZone("EST", -5*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfDay(tt.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfDay(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want time.Time
	}{
		{
			"returns the date with the time set to 00:00:00",
			time.Date(2024, 7, 5, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 7, 5, 23, 59, 59, 999, time.UTC),
		},
		{
			"keeps the original time zone",
			time.Date(2024, 7, 5, 23, 7, 8, 5, time.FixedZone("EST", -5*3600)),
			time.Date(2024, 7, 5, 23, 59, 59, 999, time.FixedZone("EST", -5*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfDay(tt.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
