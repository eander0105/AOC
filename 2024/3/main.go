package main

import (
	"os"
	"fmt"
	"regexp"
	"strconv"
)

func runMul(input string) int {
	r, _ := regexp.Compile("[0-9]{1,3}")
	values := r.FindAllString(input, -1)
	total := 0
	for _, v := range values {
		number, _ := strconv.ParseInt(v, 10, 64)
		if total > 0 {
			total *= int(number)
		} else {
			total = int(number)
		}
	}
	return total
}

func part1(input string) int {
	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	data := r.FindAllString(input, -1)

	total := 0
	for _, item := range data {
		total += runMul(item)
	}
	return total
}

func part2(input string) int {
	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)|do\\(\\)|don\\'t\\(\\)")
	data := r.FindAllString(input, -1)

	total := 0
	enabled := true
	for _, item := range data {
		if item == "do()" {
			enabled = true
		} else if item == "don't()" {
			enabled = false
		}

		if enabled {
			total += runMul(item)
		}
	}
	return total
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("no file!")
		return
	}
	fmt.Println("Result part1:", part1(string(data)))
	fmt.Println("Result part2:", part2(string(data)))
}
