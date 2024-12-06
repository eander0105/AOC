package main

import (
	"fmt"
	"os"
	"strings"
	"slices"
	"strconv"
	"math"
)

func easy(data []byte) {

	value := string(data)
	lbl := strings.Split(value, "\n")

	var left []string
	var right []string

	for _, l := range lbl {
		pair := strings.Split(l, "   ")

		left = append(left, pair[0])
		right = append(right, pair[1])
	}

	slices.Sort(left)
	slices.Sort(right)


	var distance float64 = 0
	for index, _ := range left {
		leftv, _ := strconv.ParseInt(left[index], 10, 64)
		rightv, _ := strconv.ParseInt(right[index], 10, 64)
		dst := math.Abs(float64(leftv) - float64(rightv))
		distance = distance + dst
	}
	fmt.Println(int(distance))
}

func hard(data []byte) {
	value := string(data)
	lbl := strings.Split(value, "\n")

	rightAmount := make(map[string]int)

	for _, l := range lbl {
		pair := strings.Split(l, "   ")
		rightAmount[pair[1]]++
	}


	var amount int = 0
	for _, l := range lbl {
		pair := strings.Split(l, "   ")
		if rightAmount[pair[0]] > 0 {
			v, _ := strconv.ParseInt(pair[0], 10, 64)
			amount = amount + (rightAmount[pair[0]] * int(v))
		}
	}
	fmt.Println(amount)

}

func main() {
	data, err := os.ReadFile("2.txt")
	if err != nil {
		fmt.Println("Stopid")
		return
	}
	easy(data)
	hard(data)
}
