package main

// code starts here
func Part1(input [][]byte) int {
	total := 0
	isDigit := func(r byte) bool {
		return '0' <= r && r <= '9'
	}
	for _, v := range input {
		left, right := false, false
		i, j := 0, len(v)-1
		temp := 0
		for !(left && right) {
			if !left && isDigit(v[i]) {
				temp += int(v[i]-'0') * 10
				left = true
			}
			if !right && isDigit(v[j]) {
				temp += int(v[j] - '0')
				right = true
			}
			i++
			j--
		}
		total += temp
	}
	return total
}
