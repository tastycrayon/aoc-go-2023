package main

import (
	"bytes"
	"strconv"
)

// code starts here
func Part2(input [][]byte) int {
	colon, semicolon, comma, space := []byte(": "), []byte("; "), []byte(", "), []byte(" ")
	total := 0
	getPower := func(sets [][]byte) int {
		var maxR, maxG, maxB int = 0, 0, 0
		for _, set := range sets {
			itemsInSet := bytes.Split(set, comma)
			for _, item := range itemsInSet {
				digits, color, _ := bytes.Cut(item, space)
				num, _ := strconv.Atoi(string(digits))
				switch color[0] {
				case 'r':
					if num > maxR {
						maxR = num
					}
				case 'g':
					if num > maxG {
						maxG = num
					}
				case 'b':
					if num > maxB {
						maxB = num
					}
				}
			}
		}
		return maxR * maxG * maxB
	}
	for _, v := range input {
		_, game, _ := bytes.Cut(v, colon)
		sets := bytes.Split(game, semicolon)
		total += getPower(sets)
	}

	return total
}
