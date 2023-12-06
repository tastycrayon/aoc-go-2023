package main

import (
	"regexp"
	"strings"
)

// code starts here
func Part2(input [][]byte) int {
	total := 0
	numberMap := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	reverseString := func(s []byte) string {
		sb := strings.Builder{}
		sb.Grow(len(s))
		maxIdx := len(s) - 1
		for i := range s {
			sb.WriteByte(s[maxIdx-i])
		}
		return sb.String()
	}
	reverseByte := func(s []byte) {
		i, j := 0, len(s)-1
		for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	reg := "[0-9]|one|two|three|four|five|six|seven|eight|nine"
	regInReverse := "[0-9]|" + reverseString([]byte("one|two|three|four|five|six|seven|eight|nine"))

	for _, v := range input {
		resultFirst := regexp.MustCompile(reg).FindIndex(v)
		first := string(v[resultFirst[0]:resultFirst[1]])

		reverseByte(v)
		resultLast := regexp.MustCompile(regInReverse).FindIndex(v)
		temp := v[resultLast[0]:resultLast[1]]
		reverseByte(temp)
		last := string(temp)

		var num1, num2 int
		if len(first) > 1 {
			num1 = numberMap[first]
		} else {
			num1 = int(first[0] - '0')
		}
		if len(last) > 1 {
			num2 = numberMap[last]
		} else {
			num2 = int(last[0] - '0')
		}

		total += num1*10 + num2
	}
	return total
}
