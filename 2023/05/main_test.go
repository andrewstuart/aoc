package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAOC(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)

	f, err := os.OpenFile("./test.txt", os.O_RDONLY, 0400)
	rq.NoError(err)

	out := aoc(bufio.NewReader(f))

	asrt.Equal(46, out)
}

func TestOverlaps(t *testing.T) {
	tests := []struct {
		name string
		in   [2]Range
		out  bool
	}{
		{"1-10,5-6", [2]Range{{1, 10}, {5, 2}}, true},
		{"5-6,6-7", [2]Range{{5, 2}, {6, 2}}, true},
		{"5-6,7-8", [2]Range{{5, 2}, {7, 2}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in[0].Overlaps(tt.in[1]); got != tt.out {
				t.Errorf("got %v, want %v", got, tt.out)
			}
			if got := tt.in[1].Overlaps(tt.in[0]); got != tt.out {
				t.Errorf("got %v backwards, want %v", got, tt.out)
			}
		})
	}
}

func TestMapRange(t *testing.T) {
	asrt := assert.New(t)
	m := Map{Range{5, 2}, 0} // 5-6->0-2
	r := Range{4, 4}         //4-7

	asrt.Equal([]Range{{0, 2}, {4, 1}, {7, 1}}, m.MapRange(r))
	asrt.Equal([]Range{{0, 2}}, m.MapRange(Range{5, 2}))
}

// func TestSplit(t *testing.T) {
// 	asrt := assert.New(t)

// 	m1 := Map{Range{0, 10}, 10} // 0-9 -> 10-19
// 	m2 := Map{Range{10, 5}, 20} // 10-14 -> 20-24

// 	// 0-4 -> 20-24, 5-9 -> 15-19, 10-14 -> 20-24
// 	asrt.Equal([]Map{{Range{0, 5}, 20}, {Range{5, 5}, 15}, {Range{10, 5}, 20}}, m1.Split(m2))

// 	m1 = Map{Range{0, 10}, 10} // 0-9 -> 10-19
// 	m2 = Map{Range{20, 5}, 30} // 10-14 -> 20-24
// 	asrt.Equal([]Map{m1, m2}, m1.Split(m2))
// }
