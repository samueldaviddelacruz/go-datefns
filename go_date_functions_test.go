package go_date_functions

import (
	"reflect"
	"testing"
	"time"
)

type args struct {
	dirtyDate time.Time
	amount    int
}
type testCase struct {
	name string
	args args
	want time.Time
}

func TestAddDays(t *testing.T) {
	var tests []testCase
	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests = append(tests, testCase{
		"adds the given number of days",
		args{dirtyDate: testDate, amount: 1},
		time.Date(1991, 9, 27, 0, 0, 0, 0, time.UTC),
	}, testCase{
		"subtract the given number of days",
		args{dirtyDate: testDate, amount: -1},
		time.Date(1991, 9, 25, 0, 0, 0, 0, time.UTC),
	})

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

	tests := []testCase{
		{
			"adds the given number of business days",
			args{dirtyDate: time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC), amount: 10},
			time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"subtract the given number of business days",
			args{dirtyDate: time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC), amount: -1},
			time.Date(2022, 6, 17, 0, 0, 0, 0, time.UTC),
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
	var tests []testCase
	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests = append(tests, testCase{
		"adds the given number of hours",
		args{dirtyDate: testDate, amount: 1},
		time.Date(1991, 9, 26, 1, 0, 0, 0, time.UTC),
	}, testCase{
		"subtract the given number of hours",
		args{dirtyDate: time.Date(1991, 9, 26, 1, 0, 0, 0, time.UTC), amount: -1},
		time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
	})

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
	var tests []testCase
	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests = append(tests, testCase{
		"adds the given number of minutes",
		args{dirtyDate: testDate, amount: 1},
		time.Date(1991, 9, 26, 0, 1, 0, 0, time.UTC),
	}, testCase{
		"subtract the given number of minutes",
		args{dirtyDate: time.Date(1991, 9, 26, 0, 1, 0, 0, time.UTC), amount: -1},
		time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMinutes(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMinutes() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddHours(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddMinutes() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})
}

func TestAddMilliseconds(t *testing.T) {
	var tests []testCase
	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests = append(tests, testCase{
		"adds the given number of milliseconds",
		args{dirtyDate: testDate, amount: 1000},
		time.Date(1991, 9, 26, 0, 0, 1, 0, time.UTC),
	}, testCase{
		"subtract the given number of milliseconds",
		args{dirtyDate: time.Date(1991, 9, 26, 0, 0, 1, 0, time.UTC), amount: -1000},
		time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC),
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMilliseconds(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMilliseconds() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddHours(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddMilliseconds() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})
}

func TestAddMonths(t *testing.T) {
	var tests []testCase
	testDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	tests = append(tests, testCase{
		"adds the given number of Months",
		args{dirtyDate: testDate, amount: 1},
		time.Date(1991, 10, 26, 0, 0, 0, 0, time.UTC),
	}, testCase{
		"subtract the given number of Months",
		args{dirtyDate: time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC), amount: -1},
		time.Date(1991, 8, 26, 0, 0, 0, 0, time.UTC),
	},
		testCase{
			"adds the given number of Months (February)",
			args{dirtyDate: time.Date(2022, 1, 31, 0, 0, 0, 0, time.UTC), amount: 1},
			time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC),
		},
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMonths(tt.args.dirtyDate, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMonths() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("does not mutate original date", func(t *testing.T) {
		if got := AddHours(testDate, 0); !reflect.DeepEqual(got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("AddMonths() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})
}
