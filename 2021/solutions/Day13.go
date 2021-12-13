package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const SIZE = 1500

func foldx(paper [][]bool, d, rmost, bmost int) int {
	from := 0
	to := rmost
	if d >= rmost-d {
		from = 2*d - rmost
	} else {
		to = 2 * d
	}

	for from < to {
		for i := 0; i <= bmost; i++ {
			paper[i][from] = paper[i][from] || paper[i][to]
			if paper[i][to] {
				paper[i][to] = false
			}
		}
		from++
		to--
	}
	return to
}

func foldy(paper [][]bool, d, rmost, bmost int) int {
	from := 0
	to := bmost
	if d >= bmost-d {
		from = 2*d - bmost
	} else {
		to = 2 * d
	}

	for from < to {
		for i := 0; i <= rmost; i++ {
			paper[from][i] = paper[from][i] || paper[to][i]
			if paper[to][i] {
				paper[to][i] = false
			}
		}
		from++
		to--
	}
	return to
}

func count(paper [][]bool, r, b int) int {
	cnt := 0
	for i := 0; i <= b; i++ {
		for j := 0; j <= r; j++ {
			if paper[i][j] {
				cnt++
			}
		}
	}
	return cnt
}

func D13Part1(filepath string) int {
	paper := make([][]bool, SIZE)
	for i := 0; i < SIZE; i++ {
		paper[i] = make([]bool, SIZE)
	}
	lines := readLines(filepath)
	actionsIdx := 0
	maxX := 0
	maxY := 0
	for idx, line := range lines {
		if line == "" {
			actionsIdx = idx + 1
			break
		}
		sp := strings.Split(line, ",")
		a, _ := strconv.Atoi(sp[0])
		if a > maxX {
			maxX = a
		}
		b, _ := strconv.Atoi(sp[1])
		if b > maxY {
			maxY = b
		}
		paper[b][a] = true
	}
	for i := actionsIdx; i < len(lines); i++ {
		sp := strings.Split(lines[i], "=")
		a := sp[0][len(sp[0])-1]
		d, _ := strconv.Atoi(sp[1])
		if a == 'x' {
			_ = foldx(paper, d, maxX, maxY)
		} else {
			_ = foldy(paper, d, maxX, maxY)
		}
		return count(paper, maxX, maxY)
	}
	return -1
}

func D13Part2(filepath string) {
	paper := make([][]bool, SIZE)
	for i := 0; i < SIZE; i++ {
		paper[i] = make([]bool, SIZE)
	}
	lines := readLines(filepath)
	actionsIdx := 0
	maxX := 0
	maxY := 0
	for idx, line := range lines {
		if line == "" {
			actionsIdx = idx + 1
			break
		}
		sp := strings.Split(line, ",")
		a, _ := strconv.Atoi(sp[0])
		if a > maxX {
			maxX = a
		}
		b, _ := strconv.Atoi(sp[1])
		if b > maxY {
			maxY = b
		}
		paper[b][a] = true
	}
	x := 0
	y := 0
	for i := actionsIdx; i < len(lines); i++ {
		sp := strings.Split(lines[i], "=")
		a := sp[0][len(sp[0])-1]
		d, _ := strconv.Atoi(sp[1])
		if a == 'x' {
			x = foldx(paper, d, maxX, maxY)
		} else {
			y = foldy(paper, d, maxX, maxY)
		}
	}
	f, _ := os.Create("output.txt")
	defer f.Close()
	bufwrite := bufio.NewWriter(f)
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			if paper[i][j] {
				bufwrite.WriteByte('#')
			} else {
				bufwrite.WriteByte('.')
			}
		}
		bufwrite.WriteByte('\n')
	}
	bufwrite.Flush()
}
