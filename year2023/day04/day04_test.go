package day04

import (
	"testing"

	"gotest.tools/assert"
)

func Test_lineToCard(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Card
	}{
		{
			name: "Example",
			args: args{
				line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			},
			want: Card{
				Id:             1,
				WinningNumbers: []int{41, 48, 83, 86, 17},
				MyNumbers:      []int{83, 86, 6, 31, 17, 9, 48, 53},
				Copies:         1,
			},
		},
		{
			name: "Example",
			args: args{
				line: "Card   1: 91 73 74 57 24 99 31 70 60  8 | 89 70 43 24 62 30 91 87 60 57 90  2 27  3 31 25 39 83 64 73 99  8 74 37 49",
			},
			want: Card{
				Id:             1,
				WinningNumbers: []int{91, 73, 74, 57, 24, 99, 31, 70, 60, 8},
				MyNumbers:      []int{89, 70, 43, 24, 62, 30, 91, 87, 60, 57, 90, 2, 27, 3, 31, 25, 39, 83, 64, 73, 99, 8, 74, 37, 49},
				Copies:         1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, tt.want, lineToCard(tt.args.line))
		})
	}
}

func TestCard_Points(t *testing.T) {
	type fields struct {
		Id             int
		WinningNumbers []int
		MyNumbers      []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Example",
			fields: fields{
				Id:             1,
				WinningNumbers: []int{41, 48, 83, 86, 17},
				MyNumbers:      []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			want: 4,
		},
		{
			name: "Example2",
			fields: fields{
				Id:             2,
				WinningNumbers: []int{13, 32, 20, 16, 61},
				MyNumbers:      []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Id:             tt.fields.Id,
				WinningNumbers: tt.fields.WinningNumbers,
				MyNumbers:      tt.fields.MyNumbers,
			}
			assert.Equal(t, tt.want, c.Points())
		})
	}
}

func Test_processCardCopying(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				cards: []Card{
					{
						Id:             1,
						WinningNumbers: []int{41, 48, 83, 86, 17},
						MyNumbers:      []int{83, 86, 6, 31, 17, 9, 48, 53},
						Copies:         1,
					},
					{
						Id:             2,
						WinningNumbers: []int{13, 32, 20, 16, 61},
						MyNumbers:      []int{61, 30, 68, 82, 17, 32, 24, 19},
						Copies:         1,
					},
					{
						Id:             3,
						WinningNumbers: []int{1, 21, 53, 59, 44},
						MyNumbers:      []int{69, 82, 63, 72, 16, 21, 14, 1},
						Copies:         1,
					},
					{
						Id:             4,
						WinningNumbers: []int{41, 92, 73, 84, 69},
						MyNumbers:      []int{59, 84, 76, 51, 58, 5, 54, 83},
						Copies:         1,
					},
					{
						Id:             5,
						WinningNumbers: []int{87, 83, 26, 28, 32},
						MyNumbers:      []int{88, 30, 70, 12, 93, 22, 82, 36},
						Copies:         1,
					},
					{
						Id:             6,
						WinningNumbers: []int{31, 18, 13, 56, 72},
						MyNumbers:      []int{74, 77, 10, 23, 35, 67, 36, 11},
						Copies:         1,
					},
				},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, processCardCopying(tt.args.cards))
		})
	}
}

func Test_solution(t *testing.T) {
	assert.Equal(t, 20855, solutionA())
	assert.Equal(t, 5489600, solutionB())
}
