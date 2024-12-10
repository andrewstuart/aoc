package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

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

type file struct {
	id     int
	blocks int
	start  int
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, ezaoc.IntSlicer(""))
	if err != nil {
		log.Fatal(err)
	}
	ins := inputs[0]

	count := 0
	files := map[int]*file{}
	var blocks []int
	for i, in := range ins {
		if i%2 == 0 {
			id := i / 2
			files[id] = &file{id: id, blocks: in, start: len(blocks)}
			blk := make([]int, in)
			fill(blk, id)
			blocks = append(blocks, blk...)
			continue
		}
		blk := make([]int, in)
		fill(blk, -1)
		blocks = append(blocks, blk...)
	}

	// print(blocks)
	blks := make([]int, len(blocks))
	copy(blks, blocks)

	for nextID := len(files) - 1; nextID > 0; nextID-- {
		// print(blks)
		size := files[nextID].blocks
		i := findFree(blks, 0, size)
		// fmt.Println(i, nextID, size)
		if i == -1 {
			continue
		}
		if i > files[nextID].start {
			continue
		}

		for ii := 0; ii < size; ii++ {
			blks[i+ii], blks[files[nextID].start+ii] = nextID, -1
		}

		// found := false
		// find the largest file ID with a small enough size to fit
		// for jj := nextID; jj > 0; jj-- {
		// 	f, ok := files[jj]
		// 	if !ok || f.blocks > free || f.start <= i {
		// 		continue
		// 	}
		// 	found = true
		// 	for ii := 0; ii < f.blocks; ii++ {
		// 		blks[i+ii] = f.id
		// 		blks[f.start+ii] = -1
		// 	}
		// 	i += f.blocks
		// 	delete(files, f.id)
		// 	break
		// }
		// if !found {
		// 	i++
		// }
	}

	for i, id := range blks {
		if id == -1 {
			continue
		}
		count += i * id
	}

	return count
}

func findFree(blocks []int, start, min int) int {
search:
	for i := start; i < len(blocks); i++ {
		if blocks[i] != -1 {
			continue
		}
		for j := 0; i+j < len(blocks); j++ {
			// fmt.Println("search", i, j, min)
			if blocks[i+j] != -1 {
				continue search
			}
			if j >= min-1 {
				return i
			}
		}
	}
	return -1
}

func print(is []int) {
	for _, i := range is {
		if i == -1 {
			fmt.Print(".")
			continue
		}
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func fill(bs []int, fill int) {
	for i := 0; i < len(bs); i++ {
		bs[i] = fill
	}
}
