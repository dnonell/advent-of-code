package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var elves_calories []int

	var calorie int
	var calories = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves_calories = append(elves_calories, calories)
			calories = 0
		} else {
			calorie, _ = strconv.Atoi(line)
			calories += calorie
		}
	}

	sort.Slice(elves_calories, func(i, j int) bool {
		return elves_calories[i] > elves_calories[j]
	})

	fmt.Printf("Part 1 -> %d\n", elves_calories[0])
	fmt.Printf("Part 2 -> %d\n", elves_calories[0]+elves_calories[1]+elves_calories[2])
}
