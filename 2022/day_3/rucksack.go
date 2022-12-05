package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func priorityOf(value string) int {
	abc := "abcdefghijklmnopqrstuvwxyz"
	priorities := abc + strings.ToUpper(abc)

	return strings.Index(priorities, value) + 1
}

func guessSharedItems(item_1, item_2 string) string {
	m := func(r rune) rune {
		if strings.ContainsRune(item_2, r) {
			return r
		}
		return -1
	}

	return strings.Map(m, item_1)
}

func findItemType(items []string) string {
	if len(items) == 2 {
		return guessSharedItems(items[0], items[1])[0:1]
	}

	shared_items := guessSharedItems(items[0], items[1])

	if len(shared_items) == 1 {
		return shared_items
	}

	return guessSharedItems(shared_items, items[2])[0:1]
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Part 1 -> 7903
	priorities_1 := 0
	for _, line := range lines {
		item_len := len(line) / 2
		item_type := findItemType([]string{line[0:item_len], line[item_len:]})
		priorities_1 += priorityOf(item_type)
	}
	fmt.Printf("Part 1 -> %d\n", priorities_1)

	// Part 2
	priorities_2 := 0
	group_size := 3
	total_groups := len(lines) / group_size
	for i := 0; i < total_groups; i++ {
		current_index := i * group_size
		item_type := findItemType(lines[current_index : current_index+group_size])
		priorities_2 += priorityOf(item_type)
	}
	fmt.Printf("Part 2 -> %d\n", priorities_2)
}
