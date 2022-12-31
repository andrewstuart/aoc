package main

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAll(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal("CMZ", aoc(strings.NewReader(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)))

}
