package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAOC(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)

	f, err := os.OpenFile("./test.txt", os.O_RDONLY, 0400)
	rq.NoError(err)

	out := aoc(bufio.NewReader(f), false)

	asrt.Equal(6, out)
}

func TestInput(t *testing.T) {
	asrt := assert.New(t)

	const edge = `.##..
....#
.....
.^.#.
.....`

	out := aoc(strings.NewReader(edge), false)

	asrt.Equal(1, out)
}
