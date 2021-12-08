package solutions

import (
	"strings"
)

func processLine(line string) ([]string, []string) {
	sp := strings.Split(line, " | ")

	return strings.Split(sp[0], " "), strings.Split(sp[1], " ")
}

func D08Part1(input string) int {
	lines := readInput(input)

	cnt := 0
	for _, line := range lines {
		_, p2 := processLine(line)
		for _, p := range p2 {
			if len(p) == 2 || len(p) == 3 || len(p) == 7 || len(p) == 4 {
				cnt++
			}
		}
	}
	return cnt
}

func digitStrToBit(d string) (r uint8) {
	for i := 0; i < len(d); i++ {
		r |= 1 << (byte(d[i] - byte('a')))
	}
	return
}

func ones(i uint8) int {
	r := 0
	for i != 0 {
		if (i & 1) == 1 {
			r++
		}
		i >>= 1
	}
	return r
}

func deduce(knowns []string) (deduced map[uint8]int) {
	deduced = make(map[uint8]int)
	t := make([]uint8, len(knowns))
	for i := 0; i < len(t); i++ {
		t[i] = digitStrToBit(knowns[i])
	}
	for i := 0; i < len(t); i++ {
		switch len(knowns[i]) {
		case 2:
			deduced[t[i]] = 1
		case 3:
			deduced[t[i]] = 7
		case 4:
			deduced[t[i]] = 4
		case 7:
			deduced[t[i]] = 8
		}
	}

	rev := make(map[int]uint8)
	for k, v := range deduced {
		rev[v] = k
	}

	for i := 0; i < len(t); i++ {
		switch len(knowns[i]) {
		case 5: // 2, 3, 5
			if (t[i] & rev[1]) == rev[1] {
				deduced[t[i]] = 3
			} else if ones(t[i]&rev[4]) == 3 {
				deduced[t[i]] = 5
			} else {
				deduced[t[i]] = 2
			}
		case 6: // 0, 6, 9
			if ones(t[i]&rev[4]) == 4 {
				deduced[t[i]] = 9
			} else if t[i]&rev[1] == rev[1] {
				deduced[t[i]] = 0
			} else {
				deduced[t[i]] = 6
			}
		default:
			continue
		}
	}

	return
}

func singleLine(line string) int {
	p1, p2 := processLine(line)
	// 处理前半部分
	mapping := deduce(p1)

	sum := 0
	for _, v := range p2 {
		sum = sum*10 + mapping[digitStrToBit(v)]
	}
	return sum
}

func D08Part2(input string) int {
	lines := readInput(input)

	sum := 0
	for _, line := range lines {
		sum += singleLine(line)
	}
	return sum
}
