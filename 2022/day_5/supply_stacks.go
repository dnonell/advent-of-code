package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseMove(line string) (int, int, int) {
	move := strings.Split(line, " ")
	number, _ := strconv.Atoi(move[1])
	from, _ := strconv.Atoi(move[3])
	to, _ := strconv.Atoi(move[5])
	return number, from - 1, to - 1
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	stacks_part_1 := [][]string{
		[]string{"D", "M", "S", "Z", "R", "F", "W", "N"},
		[]string{"W", "P", "Q", "G", "S"},
		[]string{"W", "R", "V", "Q", "F", "N", "J", "C"},
		[]string{"F", "Z", "P", "C", "G", "D", "L"},
		[]string{"T", "P", "S"},
		[]string{"H", "D", "F", "W", "R", "L"},
		[]string{"Z", "N", "D", "C"},
		[]string{"W", "N", "R", "F", "V", "S", "J", "Q"},
		[]string{"R", "M", "S", "G", "Z", "W", "V"},
	}
	stacks_part_2 := [][]string{
		[]string{"D", "M", "S", "Z", "R", "F", "W", "N"},
		[]string{"W", "P", "Q", "G", "S"},
		[]string{"W", "R", "V", "Q", "F", "N", "J", "C"},
		[]string{"F", "Z", "P", "C", "G", "D", "L"},
		[]string{"T", "P", "S"},
		[]string{"H", "D", "F", "W", "R", "L"},
		[]string{"Z", "N", "D", "C"},
		[]string{"W", "N", "R", "F", "V", "S", "J", "Q"},
		[]string{"R", "M", "S", "G", "Z", "W", "V"},
	}

	for index, line := range lines {
		if index < 10 {
			continue
		}

		number, from, to := parseMove(line)

		// Part 1
		from_stack_1 := stacks_part_1[from]
		from_index_1 := len(from_stack_1) - number
		stacks_part_1[from] = from_stack_1[0:from_index_1]
		for i := 1; i <= number; i++ {
			stacks_part_1[to] = append(stacks_part_1[to], from_stack_1[len(from_stack_1)-i])
		}

		// Part 2
		from_stack_2 := stacks_part_2[from]
		from_index_2 := len(from_stack_2) - number
		stacks_part_2[from] = from_stack_2[0:from_index_2]
		crates := from_stack_2[from_index_2:]
		for _, crate := range crates {
			stacks_part_2[to] = append(stacks_part_2[to], crate)
		}

	}

	fmt.Printf("FWNSHLDNZ -> %s\n", stacks_part_1)
	fmt.Printf("RNRGDNFQG -> %s\n", stacks_part_2)
}
