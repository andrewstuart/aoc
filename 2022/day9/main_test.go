package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal(13, aoc(strings.NewReader(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)))
}
