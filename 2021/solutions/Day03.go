package solutions

import (
	"log"
	"strconv"
)

func D3Part1(input string) int {
	lines := readInput(input)
	if len(lines) == 0 {
		return 0
	}
	cnt := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				cnt[i] += 1
			}
		}
	}

	gammaString := ""
	epsilonString := ""
	for _, n := range cnt {
		if n > (len(lines) / 2) {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}
	gamma, err := strconv.ParseInt(gammaString, 2, len(lines[0])+1)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(epsilonString, 2, len(lines[0])+1)
	if err != nil {
		log.Fatal(err)
	}

	return int(gamma) * int(epsilon)
}

func find(data []string, f func(int, int) bool) (ret int64) {
	validCnt := len(data)
	validIdx := make([]bool, len(data))
	for i := 0; i < len(data); i++ {
		validIdx[i] = true
	}
	pos := 0
	for validCnt != 1 {
		zeros := 0
		ones := 0
		for i := 0; i < len(data); i++ {
			if validIdx[i] && data[i][pos] == '1' {
				ones += 1
			} else if validIdx[i] && data[i][pos] == '0' {
				zeros += 1
			}
		}
		if f(ones, zeros) {
			for i, v := range validIdx {
				if v && data[i][pos] != '1' {
					validIdx[i] = false
					validCnt -= 1
				}
			}
		} else {
			for i, v := range validIdx {
				if v && data[i][pos] != '0' {
					validIdx[i] = false
					validCnt -= 1
				}
			}
		}
		pos += 1
	}

	for i, v := range validIdx {
		if v {
			ret, _ = strconv.ParseInt(data[i], 2, 32)
			return
		}
	}
	return
}

func D3Part2(input string) int64 {
	lines := readInput(input)

	o := find(lines, func(i1, i2 int) bool { return i1 >= i2 })
	c := find(lines, func(i1, i2 int) bool { return i1 < i2 })

	return o * c
}
