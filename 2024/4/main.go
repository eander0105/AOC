package main

import (
	"fmt"
	"os"
)

func part1(input string) int {

}

func part2(input string) int {

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
