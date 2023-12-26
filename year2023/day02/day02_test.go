package day02

import (
	"testing"

	"gotest.tools/assert"
)

func Test_gameIsValidWithBagConfiguration(t *testing.T) {
	type args struct {
		game             Game
		bagConfiguration BagConfiguration
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Game 1",
			args: args{
				game: Game{
					Id: 1,
					Sets: []Set{
						{
							Blues:  3,
							Reds:   4,
							Greens: 0,
						},
						{
							Blues:  6,
							Reds:   1,
							Greens: 2,
						},
						{
							Blues:  0,
							Reds:   0,
							Greens: 2,
						},
					},
				},
				bagConfiguration: BagConfiguration{
					Blues:  14,
					Reds:   12,
					Greens: 13,
				},
			},
			want: true,
		}, {
			name: "Game 2",
			args: args{
				game: Game{
					Id: 2,
					Sets: []Set{
						{
							Blues:  1,
							Reds:   0,
							Greens: 2,
						},
						{
							Blues:  4,
							Reds:   1,
							Greens: 3,
						},
						{
							Blues:  1,
							Reds:   0,
							Greens: 1,
						},
					},
				},
				bagConfiguration: BagConfiguration{
					Blues:  14,
					Reds:   12,
					Greens: 13,
				},
			},
			want: true,
		},
		{
			name: "Game 3",
			args: args{
				game: Game{
					Id: 3,
					Sets: []Set{
						{
							Blues:  6,
							Reds:   20,
							Greens: 8,
						},
						{
							Blues:  5,
							Reds:   4,
							Greens: 13,
						},
						{
							Blues:  0,
							Reds:   1,
							Greens: 5,
						},
					},
				},
				bagConfiguration: BagConfiguration{
					Blues:  14,
					Reds:   12,
					Greens: 13,
				},
			},
			want: false,
		},
		{
			name: "Game 4",
			args: args{
				game: Game{
					Id: 4,
					Sets: []Set{
						{
							Blues:  6,
							Reds:   3,
							Greens: 1,
						},
						{
							Blues:  0,
							Reds:   6,
							Greens: 3,
						},
						{
							Blues:  15,
							Reds:   14,
							Greens: 3,
						},
					},
				},
				bagConfiguration: BagConfiguration{
					Blues:  14,
					Reds:   12,
					Greens: 13,
				},
			},
			want: false,
		},
		{
			name: "Game 5",
			args: args{
				game: Game{
					Id: 5,
					Sets: []Set{
						{
							Blues:  1,
							Reds:   6,
							Greens: 3,
						},
						{
							Blues:  2,
							Reds:   1,
							Greens: 2,
						},
					},
				},
				bagConfiguration: BagConfiguration{
					Blues:  14,
					Reds:   12,
					Greens: 13,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, gameIsValidWithBagConfiguration(tt.args.game, tt.args.bagConfiguration))
		})
	}
}

func Test_parseLineToGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "Game 1",
			args: args{
				line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: Game{
				Id: 1,
				Sets: []Set{
					{
						Blues:  3,
						Reds:   4,
						Greens: 0,
					},
					{
						Blues:  6,
						Reds:   1,
						Greens: 2,
					},
					{
						Blues:  0,
						Reds:   0,
						Greens: 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, tt.want, parseLineToGame(tt.args.line))
		})
	}
}

func Test_minNecessaryBagConfigurationForGame(t *testing.T) {
	type args struct {
		game Game
	}
	tests := []struct {
		name string
		args args
		want BagConfiguration
	}{
		{
			name: "Game 1",
			args: args{
				game: Game{
					Id: 1,
					Sets: []Set{
						{
							Blues:  3,
							Reds:   4,
							Greens: 0,
						},
						{
							Blues:  6,
							Reds:   1,
							Greens: 2,
						},
						{
							Blues:  0,
							Reds:   0,
							Greens: 2,
						},
					},
				},
			},

			want: BagConfiguration{
				Blues:  6,
				Reds:   4,
				Greens: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, tt.want, minNecessaryBagConfigurationForGame(tt.args.game))
		})
	}
}

func Test_solution(t *testing.T) {
	assert.Equal(t, 2169, solutionA())
	assert.Equal(t, 60948, solutionB())
}
