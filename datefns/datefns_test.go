package datefns

import (
	"math"
	"reflect"
	"testing"
	"time"
)

type addAmountArg struct {
	dirtyDate time.Time
	amount    int
}
type datesDifArg struct {
	laterDate   time.Time
	earlierDate time.Time
}
type testCase[wantT any, argsT any] struct {
	name string
	args argsT
	want wantT
}

func TestAddDays(t *testing.T) {

	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests := []testCase[time.Time, addAmountArg]{{
		"adds the given number of days",
		addAmountArg{dirtyDate: testDate, amount: 1},
		time.Date(1991, 9, 27, 0, 0, 0, 0, time.UTC),
	}, {
		"subtract the given number of days",
		addAmountArg{dirtyDate: testDate, amount: -1},
		time.Date(1991, 9, 25, 0, 0, 0, 0, time.UTC),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDays(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
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

	tests := []testCase[time.Time, addAmountArg]{
		{
			"adds the given number of business days",
			addAmountArg{dirtyDate: time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC), amount: 10},
			time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"subtract the given number of business days",
			addAmountArg{dirtyDate: time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC), amount: -1},
			time.Date(2022, 6, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			"start on Saturday and add 1 business day",
			addAmountArg{dirtyDate: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), amount: 1},
			time.Date(2023, 7, 3, 0, 0, 0, 0, time.UTC), // Monday
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBusinessDays(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBusinessDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddHours(t *testing.T) {

	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests := []testCase[time.Time, addAmountArg]{
		{
			"adds the given number of hours",
			addAmountArg{dirtyDate: testDate, amount: 1},
			time.Date(1991, 9, 26, 1, 0, 0, 0, time.UTC),
		}, {
			"subtract the given number of hours",
			addAmountArg{dirtyDate: time.Date(1991, 9, 26, 1, 0, 0, 0, time.UTC), amount: -1},
			time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddHours(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
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
	tests := []testCase[time.Time, addAmountArg]{{
		"adds the given number of minutes",
		addAmountArg{dirtyDate: testDate, amount: 1},
		time.Date(1991, 9, 26, 0, 1, 0, 0, time.UTC),
	}, {
		"subtract the given number of minutes",
		addAmountArg{dirtyDate: time.Date(1991, 9, 26, 0, 1, 0, 0, time.UTC), amount: -1},
		time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMinutes(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
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
	tests := []testCase[time.Time, addAmountArg]{{
		"adds the given number of milliseconds",
		addAmountArg{dirtyDate: testDate, amount: 1000},
		time.Date(1991, 9, 26, 0, 0, 1, 0, time.UTC),
	}, {
		"subtract the given number of milliseconds",
		addAmountArg{dirtyDate: time.Date(1991, 9, 26, 0, 0, 1, 0, time.UTC), amount: -1000},
		time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMilliseconds(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
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
	tests := []testCase[time.Time, addAmountArg]{
		{
			"adds the given number of Months",
			addAmountArg{dirtyDate: testDate, amount: 1},
			time.Date(1991, 10, 26, 0, 0, 0, 0, time.UTC),
		}, {
			"subtract the given number of Months",
			addAmountArg{dirtyDate: time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC), amount: -1},
			time.Date(1991, 8, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			"adds the given number of Months (February)",
			addAmountArg{dirtyDate: time.Date(2022, 1, 31, 0, 0, 0, 0, time.UTC), amount: 1},
			time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			"adds month from Jan 31 to leap Feb",
			addAmountArg{dirtyDate: time.Date(2020, 1, 31, 0, 0, 0, 0, time.UTC), amount: 1},
			time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMonths(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
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
	tests := []testCase[bool, time.Time]{
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
			if got := IsWeekend(tt.args); got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInCalendarDays(t *testing.T) {
	tests := []testCase[int, datesDifArg]{
		{
			"returns the number of full days between the given dates",
			datesDifArg{laterDate: time.Date(2012, 7, 2, 18, 0, 0, 0, time.UTC), earlierDate: time.Date(2011, 7, 2, 6, 0, 0, 0, time.UTC)},
			366,
		},
		{
			"returns no full days between the given dates",
			datesDifArg{laterDate: time.Date(2024, 7, 6, 5, 0, 0, 0, time.UTC), earlierDate: time.Date(2024, 7, 6, 4, 59, 0, 0, time.UTC)},
			0,
		},
		{
			"returns the number of full days between the given dates in negative",
			datesDifArg{earlierDate: time.Date(2012, 7, 2, 18, 0, 0, 0, time.UTC), laterDate: time.Date(2011, 7, 2, 6, 0, 0, 0, time.UTC)},
			-366,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceInCalendarDays(tt.args.laterDate, tt.args.earlierDate); got != tt.want {
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
