package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
)

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br))
}

type cmd struct {
	Input  string
	Output []string
}

type entry struct {
	name string
	size int
	dir  map[string]*entry
}

func (e entry) du() int {
	if e.size > 0 {
		return e.size
	}
	size := 0
	for _, e := range e.dir {
		size += e.du()
	}
	return size
}

func (e entry) walk(f func(*entry)) {
	f(&e)
	if e.dir == nil {
		return
	}
	for _, e := range e.dir {
		e.walk(f)
	}
}

func (e entry) get(s ...string) *entry {
	var next = &e
	for _, s := range s {
		if dir, ok := next.dir[s]; ok {
			next = dir
			continue
		}
		next.dir[s] = newEntry(s)
		next = next.dir[s]
	}
	return next
}

func newEntry(n string) *entry {
	return &entry{
		name: n,
		dir:  map[string]*entry{},
	}
}

func aoc(r io.Reader) int {
	var in cmd
	inputs, err := ezaoc.ReadAOC(r, func(st string) (cmd, error) {
		if st == "" {
			return cmd{}, io.EOF
		}
		if strings.HasPrefix(st, "$") {
			i := in
			in = cmd{Input: strings.TrimPrefix(st, "$ ")}
			return i, nil
		}
		in.Output = append(in.Output, st)
		return in, ezaoc.ErrIgnore
	})
	if err != nil {
		log.Fatal(err)
	}
	inputs = inputs[1:]

	root := newEntry("")
	count := 0
	pwd := ezaoc.Stack[string]{}
	for _, in := range inputs {
		switch strings.Fields(in.Input)[0] {
		case "ls":
			dir := root.get(pwd...)
			// if dir.dir == nil {
			// 	dir.dir = map[string]*entry{}
			// }
			for _, o := range in.Output {
				fs := strings.Fields(o)
				if n, err := strconv.Atoi(fs[0]); err == nil {
					dir.dir[fs[1]] = &entry{
						name: fs[1],
						size: n,
					}
					continue
				}
				dir.dir[fs[1]] = &entry{
					name: fs[1],
					dir:  map[string]*entry{},
				}
			}
		case "cd":
			switch strings.Fields(in.Input)[1] {
			case "..":
				pwd.Pop()
			case "/":
				pwd = ezaoc.Stack[string]{}
			default:
				pwd.Push(strings.Fields(in.Input)[1])
			}
		}
	}

	root.walk(func(e *entry) {
		if e.dir == nil {
			return
		}
		if e.du() <= 100_000 {
			count += e.du()
		}
	})

	return count
}
