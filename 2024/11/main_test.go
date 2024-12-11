package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAOC(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)

	f, err := os.ReadFile("./test.txt")
	rq.NoError(err)

	out := aoc(bytes.NewReader(f), 25)
	asrt.Equal(55312, out)
	out = aoc(bytes.NewReader(f), 6)
	asrt.Equal(22, out)
}
