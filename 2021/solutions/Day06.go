package solutions

import (
	"strconv"
	"strings"
)

func D06Part1(input string) int {
	rawData := readInput(input)
	initialStates := strings.Split(rawData[0], ",")

	states := make([]int, 0)
	for _, s := range initialStates {
		remain, _ := strconv.Atoi(s)
		states = append(states, remain)
	}

	for i := 0; i < 80; i++ {
		size := len(states)
		for j := 0; j < size; j++ {
			if states[j] == 0 {
				states = append(states, 8)
				states[j] = 6
			} else {
				states[j] -= 1
			}
		}
	}
	return len(states)
}

func D06Part2(input string) int64 {
	counts := make([]int64, 9)
	rawData := readInput(input)
	initialStates := strings.Split(rawData[0], ",")

	for _, s := range initialStates {
		remain, _ := strconv.Atoi(s)
		counts[remain] += 1
	}

	for i := 0; i < 256; i++ {
		c := counts[0]
		for j := 0; j < 8; j++ {
			counts[j] = counts[j+1]
		}
		counts[8] = c
		counts[6] += c
	}

	var cnt int64 = 0
	for i := 0; i < 9; i++ {
		cnt += counts[i]
	}

	return cnt
}
