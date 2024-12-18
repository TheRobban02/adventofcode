package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day1input.txt")

	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	fmt.Printf("Part 1: %d\n", calculateDifference(list1, list2))
	fmt.Printf("Part 2: %d\n", calculateSimilarityScore(list1, list2))
}

func calculateDifference(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	difference := 0

	for i, num1 := range list1 {
		var d = num1 - list2[i]
		if d < 0 {
			difference += d * -1
		} else {
			difference += d
		}
	}
	return difference
}

func calculateSimilarityScore(list1 []int, list2 []int) int {
	freq := make(map[int]int)
	for _, num := range list2 {
		freq[num] = freq[num] + 1
	}

	score := 0

	for _, num1 := range list1 {
		if val, exists := freq[num1]; exists {
			score += num1 * val
		}
	}

	return score
}
