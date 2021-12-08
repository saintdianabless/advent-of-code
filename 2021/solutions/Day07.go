package solutions

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func D07Part1(input string) int {
	rawData := readInput(input)
	strData := strings.Split(rawData[0], ",")

	data := make([]int, len(strData))
	for i := 0; i < len(strData); i++ {
		data[i], _ = strconv.Atoi(strData[i])
	}

	sort.Ints(data)

	minimum := 0x7FFFFFFF

	for i := data[0]; i <= data[len(data)-1]; i++ {
		total := 0
		for j := 0; j < len(data); j++ {
			total += int(math.Abs(float64(data[j] - i)))
		}
		if total < minimum {
			minimum = total
		}
	}
	return minimum
}

func D07Part2(input string) int {
	rawData := readInput(input)
	strData := strings.Split(rawData[0], ",")

	data := make([]int, len(strData))
	for i := 0; i < len(strData); i++ {
		data[i], _ = strconv.Atoi(strData[i])
	}

	sort.Ints(data)

	minimum := 0x7FFFFFFF

	memorize := make(map[int]int)

	for i := data[0]; i <= data[len(data)-1]; i++ {
		total := 0
		for j := 0; j < len(data); j++ {
			diff := int(math.Abs(float64(data[j] - i)))
			if v, ok := memorize[diff]; ok {
				total += v
			} else {
				costs := 0
				for x := 1; x <= diff; x++ {
					costs += x
				}
				memorize[diff] = costs
				total += costs
			}
		}
		if total < minimum {
			minimum = total
		}
	}
	return minimum
}
