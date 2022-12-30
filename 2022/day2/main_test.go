package main

import "testing"

func TestRPS(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    Strat
	}{
		{"winrock", 1 + Win + int(Paper), Strat{Rock, Win}},
		{"loserock", 1 + Loss + int(Scissors), Strat{Rock, Loss}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.score()
			if actual != tt.expected {
				t.Errorf("(%+v): expected %d, actual %d: played %d", tt.given, tt.expected, actual, tt.given.me())
			}

		})
	}
}
