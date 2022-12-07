package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isInRange(x_1, y_1, x_2, y_2 string) bool {
	if x_1 <= x_2 && y_1 >= y_2 {
		fmt.Printf("%s-%s contains %s-%s\n", x_1, y_1, x_2, y_2)
		return true
	}

	if x_2 <= x_1 && y_2 >= y_1 {
		fmt.Printf("%s-%s contains %s-%s\n", x_2, y_2, x_1, y_1)
		return true
	}

	fmt.Printf("%s-%s and %s-%s doesn't overlaps completely\n", x_1, y_1, x_2, y_2)
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

	// Part 1
	overlaps := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		section_1, section_2 := strings.Split(sections[0], "-"), strings.Split(sections[1], "-")
		if isInRange(section_1[0], section_1[1], section_2[0], section_2[1]) {
			overlaps++
		}
	}
	fmt.Printf("Part 1 -> %d\n", overlaps)
}
