package main

import (
	"fmt"
	"os"
	"strings"
)



func checkCrossPattern(x int, y int, input []string) int {
	matches := 0
	return matches
}

func part1(input string) int {
	rows := strings.Split(input, "\n")
	total := 0
	for x, row := range rows {
		fmt.Println(row)
		for y, l := range row {
			letter := string(l)
			if letter == "X" {
				total += checkCrossPattern(x, y, rows)
			}
		}
	}
	return total
}

func part2(input string) int {
	return 0
}

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("no file!")
		return
	}
	fmt.Println("Result part1:", part1(string(data)))
	fmt.Println("Result part2:", part2(string(data)))
}
