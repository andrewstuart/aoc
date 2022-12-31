package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	asrt := assert.New(t)
	asrt.Equal(7, aoc(strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb")))
}
