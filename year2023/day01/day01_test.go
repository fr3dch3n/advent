package day01

import (
	"testing"

	"gotest.tools/assert"
)

func Test_extractFirstAndLastDigit(t *testing.T) {
	type args struct {
		line       string
		spelledOut bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"1234", false}, 14},
		{"test2", args{"1abc2", false}, 12},
		{"test3", args{"pqr3stu8vwx", false}, 38},
		{"test4", args{"a1b2c3d4e5f", false}, 15},
		{"test5", args{"treb7uchet", false}, 77},

		{"test6", args{"two1nine", true}, 29},
		{"test7", args{"eightwothree", true}, 83},
		{"test8", args{"abcone2threexyz", true}, 13},
		{"test9", args{"xtwone3four", true}, 24},
		{"test10", args{"4nineeightseven2", true}, 42},
		{"test11", args{"zoneight234", true}, 14},
		{"test12", args{"7pqrstsixteen", true}, 76},
		{"test13", args{"fourfourthreehnbhkmscqxdfksg64bvpppznkh", true}, 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractFirstAndLastDigit(tt.args.line, tt.args.spelledOut)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_subStringIsSpelledOutDigit(t *testing.T) {
	type args struct {
		line  string
		start int
		end   int
	}
	tests := []struct {
		name  string
		args  args
		idx   int
		digit int
	}{
		{
			"test1",
			args{"one", 0, 1},
			-1,
			-1,
		},
		{
			"test2",
			args{"one", 0, 3},
			0,
			1,
		},
		{
			"test3",
			args{"onefour", 2, 7},
			3,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := subStringIsSpelledOutDigit(tt.args.line, tt.args.start, tt.args.end)
			assert.Equal(t, tt.idx, got)
			assert.Equal(t, tt.digit, got1)
		})
	}
}

func Test_solution(t *testing.T) {
	assert.Equal(t, 54304, solutionA())
	assert.Equal(t, 54418, solutionB())
}
