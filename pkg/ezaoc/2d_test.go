package ezaoc

import (
	"fmt"
	"testing"
)

func TestTurn(t *testing.T) {
	tests := []struct {
		name string
		dir  Direction
		t    int
		want Direction
	}{
		{"Up Right", Up, TurnRight, Right},
		{"Up Left", Up, TurnLeft, Left},
		{"Up DiagLeft", Up, TurnDiagLeft, UpLeft},
		{"UpLeft Right", UpLeft, TurnRight, UpRight},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.Turn(tt.t); got != tt.want {
				t.Errorf("%s.Turn(%d) got %v, want %v", tt.dir, tt.t, got, tt.want)
			}
		})
	}
}

func TestCoprime(t *testing.T) {
	tests := []struct {
		in  [2]int
		out [2]int
	}{
		{[2]int{2, 3}, [2]int{2, 3}},
		{[2]int{4, 6}, [2]int{2, 3}},
		{[2]int{3, 9}, [2]int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.in[0], tt.in[1]), func(t *testing.T) {
			if i, j := ReduceToCoprime(tt.in[0], tt.in[1]); [2]int{i, j} != tt.out {
				t.Errorf("got %v, want %v", [2]int{i, j}, tt.out)
			}
		})
	}
}
