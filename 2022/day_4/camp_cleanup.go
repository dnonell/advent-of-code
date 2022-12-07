package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isFullyInRange(x_1, y_1, x_2, y_2 int) bool {
	if x_1 <= x_2 && y_1 >= y_2 {
		return true
	}

	if x_2 <= x_1 && y_2 >= y_1 {
		return true
	}

	return false
}

func isPartialInRange(x_1, y_1, x_2, y_2 int) bool {
	if isFullyInRange(x_1, y_1, x_2, y_2) {
		return true
	}

	if x_1 <= x_2 && y_1 >= x_2 {
		return true
	}

	if x_2 <= x_1 && y_2 >= x_1 {
		return true
	}

	return false
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part_1_overlaps, part_2_overlaps := 0, 0

	for _, line := range lines {
		sections := strings.Split(line, ",")
		section_1, section_2 := strings.Split(sections[0], "-"), strings.Split(sections[1], "-")
		x_1, _ := strconv.Atoi(section_1[0])
		y_1, _ := strconv.Atoi(section_1[1])
		x_2, _ := strconv.Atoi(section_2[0])
		y_2, _ := strconv.Atoi(section_2[1])
		if isFullyInRange(x_1, y_1, x_2, y_2) {
			part_1_overlaps++
		}
		if isPartialInRange(x_1, y_1, x_2, y_2) {
			part_2_overlaps++
		}
	}

	fmt.Printf("Part 1 -> %d\n", part_1_overlaps)
	fmt.Printf("Part 2 -> %d\n", part_2_overlaps)
}
