package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safeReportsCount = 0

	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Fields(report)

		var intArr []int
		for _, chr := range levels {
			tempNum, _ := strconv.Atoi(chr)
			intArr = append(intArr, tempNum)
		}

		var diffs []int

		for i := 0; i < len(intArr)-1; i++ {
			diffs = append(diffs, intArr[i]-intArr[i+1])
		}

		if ascendingOrDescending(diffs) {
			safeReportsCount++
		}
	}
	fmt.Printf("Part 1: %d\n", safeReportsCount)
}

func ascendingOrDescending(diffs []int) bool {
	isAscending := true
	isDescending := true

	for _, diff := range diffs {
		if diff < 1 || diff > 3 {
			isAscending = false
		}
		if diff > -1 || diff < -3 {
			isDescending = false
		}
	}

	return isAscending != isDescending
}
