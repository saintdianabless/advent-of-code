package solutions

import (
	"bufio"
	"log"
	"os"
)

func readInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			input = append(input, scanner.Text())
		}
	}
	return input
}
