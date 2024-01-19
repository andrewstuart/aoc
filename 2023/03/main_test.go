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
	// TODO replace assert expected value here
	asrt.Equal(4361, out)
}

func TestCheckRow(t *testing.T) {
	asrt := assert.New(t)
	data := []string{"....*"}
	asrt.True(checkRow(data, 0, 0, 4))

	data = []string{"*...."}
	asrt.True(checkRow(data, 0, 1, 4))
}
