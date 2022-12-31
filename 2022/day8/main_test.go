package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	asrt := assert.New(t)
	asrt.Equal(8, aoc(strings.NewReader(`30373
25512
65332
33549
35390`)))
}

func TestSplit(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal(
		[][]byte{[]byte("cba"), []byte("efg")},
		split([]byte("abcdefg"), 3),
	)
}
