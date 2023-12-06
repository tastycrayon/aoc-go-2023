package main

import "bytes"

// code starts here
func Part1(input [][]byte) int {

	total := 0
	for _, v := range input {
		removedSlice := bytes.TrimLeftFunc(v, func(r rune) bool {
			return r == ':'
		})
		winningNums, givenNums, _ := bytes.Cut(removedSlice, []byte(" | "))

	}
	return total
}
