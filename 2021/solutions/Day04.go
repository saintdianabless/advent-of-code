package solutions

import (
	"strconv"
	"strings"
)

func getNums(line string, sep string) []int {
	var numstrs []string
	if sep == " " {
		numstrs = strings.Fields(line)
	} else {
		numstrs = strings.Split(line, sep)
	}
	nums := make([]int, len(numstrs))
	for i := 0; i < len(numstrs); i++ {
		nums[i], _ = strconv.Atoi(numstrs[i])
	}
	return nums
}

func readBoard(data []string, rows, columns int) [][]int {
	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = getNums(data[i], " ")
	}
	return board
}

func readBoards(data []string, rows, columns int) [][][]int {
	boards := make([][][]int, 0)
	size := len(data)

	for idx := 0; idx < size; idx += rows {
		block := readBoard(data[idx:], rows, columns)
		boards = append(boards, block)
	}
	return boards
}

func mark(data [][]int, num int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == num {
				data[i][j] = -1
			}
		}
	}
}

func check(data [][]int) bool {
outer1:
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] != -1 {
				continue outer1
			}
		}
		return true
	}
outer2:
	for j := 0; j < len(data[0]); j++ {
		for i := 0; i < len(data); i++ {
			if data[i][j] != -1 {
				continue outer2
			}
		}
		return true
	}
	return false
}

func D04Part1(input string) int {
	rawData := readInput(input)
	nums := getNums(rawData[0], ",")

	rawData = rawData[1:]
	boards := readBoards(rawData, 5, 5)

	index1, index2 := 0, 0

outer:
	for i, num := range nums {
		for j, board := range boards {
			mark(board, num)
			if check(board) {
				index1 = i
				index2 = j
				break outer
			}
		}
	}

	m := 0
	for _, row := range boards[index2] {
		for _, num := range row {
			if num != -1 {
				m += num
			}
		}
	}

	return nums[index1] * m
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func D04Part2(input string) int {
	rawData := readInput(input)
	nums := getNums(rawData[0], ",")

	rawData = rawData[1:]
	boards := readBoards(rawData, 5, 5)

	winIndex := make([]int, 0)
	lastWin := -1
	total := len(boards)
	lastNum := -1
outer:
	for _, num := range nums {
		for j, board := range boards {
			if contains(winIndex, j) {
				continue
			}
			mark(board, num)
			if check(board) {
				winIndex = append(winIndex, j)
			}
			if len(winIndex) == total {
				lastWin = j
				lastNum = num
				break outer
			}
		}
	}

	m := 0
	for _, row := range boards[lastWin] {
		for _, num := range row {
			if num != -1 {
				m += num
			}
		}
	}

	return m * lastNum
}
