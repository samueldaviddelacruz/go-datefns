package go_date_functions

import (
	"fmt"
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

func ExampleAddDays() {
	exampleDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	exampleDatePlus1Day := AddDays(exampleDate, 1)
	fmt.Printf("testDate= %v", exampleDate)
	fmt.Printf("exampleDatePlus1Day= %v", exampleDatePlus1Day)
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

func ExampleAddBusinessDays() {
	exampleDate := time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC)
	exampleDatePlus1Day := AddBusinessDays(exampleDate, 1)
	fmt.Printf("testDate= %v", exampleDate)
	fmt.Printf("exampleDatePlus1Day= %v", exampleDatePlus1Day)
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

func ExampleAddHours() {
	exampleDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	exampleDatePlus1Hour := AddHours(exampleDate, 1)
	fmt.Printf("testDate= %v", exampleDate)
	fmt.Printf("exampleDatePlus1hour= %v", exampleDatePlus1Hour)
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
			t.Errorf("AddHours() = %v, want %v", got, time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC))
		}
	})
}

func ExampleAddMinutes() {
	exampleDate := time.Date(1991, 9, 26, 0, 0, 0, 0, time.UTC)
	exampleDatePlus1minute := AddMinutes(exampleDate, 1)
	fmt.Printf("testDate= %v", exampleDate)
	fmt.Printf("exampleDatePlus1minute= %v", exampleDatePlus1minute)
}
