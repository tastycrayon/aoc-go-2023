package main

import (
	"bytes"
	"strconv"
)

// code starts here
func Part1(input [][]byte) int {
	colon, semicolon, comma, space := []byte(": "), []byte("; "), []byte(", "), []byte(" ")
	var maxR, maxG, maxB int = 12, 13, 14
	total := 0
	checkPossibiltiy := func(sets [][]byte) bool {
		for _, set := range sets {
			itemsInSet := bytes.Split(set, comma)
			for _, item := range itemsInSet {
				digits, color, _ := bytes.Cut(item, space)
				num, _ := strconv.Atoi(string(digits))
				switch color[0] {
				case 'r':
					if num > maxR {
						return false
					}
				case 'g':
					if num > maxG {
						return false
					}
				case 'b':
					if num > maxB {
						return false
					}
				}
			}
		}
		return true
	}
	for i, v := range input {
		_, game, _ := bytes.Cut(v, colon)
		sets := bytes.Split(game, semicolon)
		if checkPossibiltiy(sets) {
			total += i + 1
		}
	}

	return total
}
