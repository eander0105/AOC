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

func safeReport(report string) bool {
	direction := '.'
	var lastValue int64 = -1

	for _, v := range strings.Fields(report) {
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
			// Check if 
			if !checkValue(direction == '+', lastValue, value, 3) {
				return false
			}
		}
		lastValue = value
	}
	return true
}

func part1(data []string) {
	safeReports := 0
	for _, row := range data {
		check := safeReport(row)
		if check {
			safeReports++
		}
	}
	fmt.Println("Safe reports:", safeReports)
}



func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("no file!")
		return
	}

	value := strings.Split(string(data), "\n")

	part1(value)
}
