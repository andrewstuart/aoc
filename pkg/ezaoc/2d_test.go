package ezaoc

import "testing"

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
