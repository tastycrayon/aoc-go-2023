package main

import (
	"strconv"
	"strings"
)

// code starts here
func Part2(input [][]byte) int {

	isDigit := func(b byte) bool {
		return '0' <= b && b <= '9'
	}
	isStar := func(b byte) bool {
		return b == '*'
	}
	maxRow, maxCol := len(input), len(input[0])
	checkPosition := func(i, j int) int {
		if i < 0 || i >= maxRow || j < 0 || j >= maxCol {
			return 0 // out of boundary index
		}
		if !isDigit(input[i][j]) {
			return 0
		}
		n, m := j, j
		for m-1 >= 0 && isDigit(input[i][m-1]) {
			m--
		}
		for n+1 < maxCol && isDigit(input[i][n+1]) {
			n++
		}
		sb := strings.Builder{}
		sb.Grow(n - m)
		for p := m; p < n+1; p++ {
			sb.WriteByte(input[i][p])
			input[i][p] = '.'
		}
		num, _ := strconv.Atoi(sb.String())
		return num
	}
	total := 0
	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			if !isStar(input[i][j]) {
				continue
			}
			top := checkPosition(i-1, j)           // top
			bottom := checkPosition(i+1, j)        // bottom
			left := checkPosition(i, j-1)          // left
			right := checkPosition(i, j+1)         // right
			bottomRight := checkPosition(i+1, j+1) // bottom-right
			bottomLeft := checkPosition(i+1, j-1)  // bottom-left
			topRight := checkPosition(i-1, j+1)    // top-right
			topLeft := checkPosition(i-1, j-1)     // top-left
			r := []int{top, bottom, left, right, bottomRight, bottomLeft, topRight, topLeft}
			gearOne, gearTwo := 0, 0
			for _, k := range r {
				if k == 0 {
					continue
				}
				if gearOne == 0 {
					gearOne = k
					continue
				}
				gearTwo = k
			}
			if gearOne == 0 || gearTwo == 0 {
				continue // one gear number missing
			}
			total += gearOne * gearTwo
		}
	}
	return total
}
