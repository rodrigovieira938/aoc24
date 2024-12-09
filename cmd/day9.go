package main

import (
	"fmt"
	"strconv"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func generate_blockmap(diskmap string) []int {
	filenum := 0
	blockmap := []int{}
	for i, number_str := range diskmap {
		number, error := strconv.Atoi(string(number_str))
		if error != nil {
			continue
		}
		for x := 0; x < number; x++ {
			if i%2 == 0 {
				blockmap = append(blockmap, filenum)
			} else {
				blockmap = append(blockmap, -1)
			}
		}
		if i%2 == 0 {
			filenum++
		}

	}
	return blockmap
}

func puzzle1(blockmap []int) int {
	for i := len(blockmap) - 1; i >= 0; i-- {
		//find first blank space
		blank_space := -1
		for x, c := range blockmap {
			//if x is bigger than i, it would fragment instead of compress
			if c == -1 && i > x {
				blank_space = x
				break
			}
		}
		if blank_space == -1 {
			break
		}
		blockmap[blank_space] = blockmap[i]
		blockmap[i] = -1
	}
	checksum := 0

	for filepos, fileid := range blockmap {
		if fileid == -1 {
			break
		}
		checksum += filepos * fileid
	}

	return checksum
}
func puzzle2(blockmap []int) int {
	for i := len(blockmap) - 1; i >= 0; i-- {
		if blockmap[i] == -1 {
			continue
		}
		fileid := blockmap[i]
		start_i := i
		size := 1
		//Get the size for this filenumber
		for {
			if i-1 >= 0 {
				if blockmap[i-1] == fileid {
					size++
					i--
				} else {
					break
				}
			} else {
				break
			}
		}

		blank_count := 0
		blank_start := -1
		for x := 0; x < start_i; x++ {
			if blockmap[x] == -1 {
				if blank_start == -1 {
					blank_start = x
				}
				blank_count++
			} else {
				blank_count = 0
				blank_start = -1
			}
			if blank_count == size {
				break
			}
		}
		if blank_start != -1 && blank_count == size {
			for x := 0; x < size; x++ {
				blockmap[start_i-x] = -1
			}
			for x := 0; x < size; x++ {
				blockmap[blank_start+x] = fileid
			}
		}
	}

	checksum := 0

	for filepos, fileid := range blockmap {
		if fileid == -1 {
			continue
		}
		checksum += filepos * fileid
	}

	return checksum
}

func main() {
	diskmap, error := utils.ReadFileStr("./data/day9.txt")
	if error == nil {
		fmt.Printf("Day 9 - Puzzle 1 answer is %v\n", puzzle1(generate_blockmap(diskmap)))
		fmt.Printf("Day 9 - Puzzle 2 answer is %v\n", puzzle2(generate_blockmap(diskmap)))
	}
}
