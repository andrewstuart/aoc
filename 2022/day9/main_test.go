package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal(36, aoc(strings.NewReader(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)))
}
