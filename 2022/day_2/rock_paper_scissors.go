package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func shapeScore(hand string) (int) {
	switch hand {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}

	fmt.Errorf("Invalid hand %s\n", hand)
	return -1
}

func convertHand(hand string) (string) {
	switch hand {
	case "X":
		return "A"
	case "Y":
		return "B"
	case "Z":
		return "C"
	}

	fmt.Errorf("Invalid hand %s\n", hand)
	return ""
}

func partOne(opponent_hand, player_hand string) (int) {
	shape_score := shapeScore(player_hand)
	
	if opponent_hand == player_hand {
		return shape_score + 3
	}

	switch opponent_hand {
	case "A":
		if player_hand == "B" {
			return shape_score + 6
		} else {
			return shape_score
		}
	case "B":
		if player_hand == "C" {
			return shape_score + 6
		} else {
			return shape_score
		}
	case "C":
		if player_hand == "A" {
			return shape_score + 6
		} else {
			return shape_score
		}
	}

	fmt.Errorf("Invalid turn -> opponent %s vs. player %s\n", opponent_hand, player_hand)
	return -1
}

func partTwo(opponent_hand, player_action string) (int) {
	var player_hand string

	switch player_action {
	case "X":
		if opponent_hand == "A" {
			return shapeScore("C")
		} else if opponent_hand == "B" {
			return shapeScore("A")
		} else {
			return shapeScore("B")
		}
	case "Y":
		player_hand = opponent_hand
		return shapeScore(player_hand) + 3
	case "Z":
		if opponent_hand == "A" {
			return shapeScore("B") + 6
		} else if opponent_hand == "B" {
			return shapeScore("C") + 6
		} else {
			return shapeScore("A") + 6
		}
	default:
		fmt.Errorf("Invalid player action -> opponent %s vs. player %s\n", opponent_hand, player_action)
		return -1
	}
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	
	part_one_score, part_two_score := 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		
		part_one_score += partOne(line[0], convertHand(line[1]))
		part_two_score += partTwo(line[0], line[1])
	}

	fmt.Printf("Part 1 -> %d\n", part_one_score)
	fmt.Printf("Part 2 -> %d\n", part_two_score)
}
