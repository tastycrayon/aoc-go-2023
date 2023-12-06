package main

import (
	"strconv"
	"strings"
)

// code starts here
func Part1(input [][]byte) int {

	isDigit := func(b byte) bool {
		return '0' <= b && b <= '9'
	}
	isDot := func(b byte) bool {
		return b == '.'
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
			item := input[i][j]
			if isDot(item) || isDigit(item) {
				continue
			}
			total += checkPosition(i-1, j)   // top
			total += checkPosition(i+1, j)   // bottom
			total += checkPosition(i, j-1)   // left
			total += checkPosition(i, j+1)   // right
			total += checkPosition(i+1, j+1) // bottom-right
			total += checkPosition(i+1, j-1) // bottom-left
			total += checkPosition(i-1, j+1) // top-right
			total += checkPosition(i-1, j-1) // top-left
		}
	}
	return total
}
