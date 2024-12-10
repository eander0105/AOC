package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func checkValue(incresing bool, lastValue int64, value int64, maxChange float64) bool {
	if value == lastValue { return false }
	if incresing {
		if value < lastValue { return false }
	} else {
		if value > lastValue { return false }
	}
	if math.Abs(float64(value) - float64(lastValue)) > maxChange {
		return false
	}
	return true
}

func safeReport(report []string) bool {
	direction := '.'
	var lastValue int64 = -1

	for _, v := range report {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return false
		}

		if lastValue >= 0 {
			// first value in report
			if direction == '.' {
				if value > lastValue {
					direction = '+'
				} else {
					direction = '-'
				}
			}
			if !checkValue(direction == '+', lastValue, value, 3) {
				return false
			}
		}
		lastValue = value
	}
	return true
}

func delete_at_index(a []string, i int) []string {
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a
}

func part1(data []string) int {
	safeReports := 0
	for _, row := range data {
		check := safeReport(strings.Fields(row))
		if check {
			safeReports++
		}
	}
	return safeReports
}

func part2(data []string) int {
	safeReports := 0
	for _, row := range data {
		report := strings.Fields(row)
		fmt.Println(report)
		for field, _ := range report {
			if field > 0 {
				r := []string{}
				copy(r[:], report)
				cReport := delete_at_index(r, field)
				fmt.Println(cReport)
				
			}

		}

		break
		// check := safeReport(row, 1)
		// fmt.Printf("Report %v, [%s] was %v\n", i ,row, check)
		// if check {
		// 	safeReports++
		// }
	}
	return safeReports
}

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("no file!")
		return
	}

	value := strings.Split(string(data), "\n")

	fmt.Println("Part1, safe reports:", part1(value))
	fmt.Println("Part2, safe reports:", part2(value))
}
