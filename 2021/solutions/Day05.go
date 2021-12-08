package solutions

import (
	"math"
	"strconv"
	"strings"
)

func maxmin(a, b int) (int, int) {
	if a >= b {
		return a, b
	} else {
		return b, a
	}
}

func markLine(mat *[1000][1000]int, x1, y1, x2, y2 int) {
	if x1 == x2 {
		max, min := maxmin(y1, y2)
		for i := min; i <= max; i++ {
			mat[x1][i] += 1
		}
	} else if y1 == y2 {
		max, min := maxmin(x1, x2)
		for i := min; i <= max; i++ {
			mat[i][y1] += 1
		}
	} else if (x1 == y2 && x2 == y1) || (math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2))) { /////   PART2 FROM HERE
		if y1 > y2 {
			x1, x2, y1, y2 = x2, x1, y2, y1
		}
		if x1 > x2 {
			for i, j := x2, y2; i <= x1; {
				mat[i][j] += 1
				i += 1
				j -= 1
			}
		} else {
			for i, j := x1, y1; i <= x2; {
				mat[i][j] += 1
				i += 1
				j += 1
			}
		}
	} else if x1 == y1 && x2 == y2 {
		if x1 < x2 {
			for i, j := x1, y1; i <= x2; {
				mat[i][j] += 1
				i += 1
				j += 1
			}
		} else {
			for i, j := x2, y2; i <= x1; {
				mat[i][j] += 1
				i += 1
				j += 1
			}
		}
	}
}

func D05Part1(input string) int {
	lines := readInput(input)
	board := [1000][1000]int{}

	for _, line := range lines {
		points := strings.Split(line, " -> ")
		t1 := strings.Split(points[0], ",")
		t2 := strings.Split(points[1], ",")
		x1, _ := strconv.Atoi(t1[0])
		y1, _ := strconv.Atoi(t1[1])
		x2, _ := strconv.Atoi(t2[0])
		y2, _ := strconv.Atoi(t2[1])
		markLine(&board, x1, y1, x2, y2)
	}

	cnt := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if board[i][j] >= 2 {
				cnt += 1
			}
		}
	}
	return cnt
}
