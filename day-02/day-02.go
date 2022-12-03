package day02

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

func decodePlay(encryptedPlay string) string {
	plays := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	result, ok := plays[encryptedPlay]
	if !ok {
		panic(fmt.Sprintf("%s is not a valid play", encryptedPlay))
	}
	return result
}

func decodeOutcome(encryptedOutcome string) int {
	outcomes := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	result, ok := outcomes[encryptedOutcome]
	if !ok {
		panic(fmt.Sprintf("%s is not a valid outcome", encryptedOutcome))
	}
	return result
}

func getWinsAndLosses() (map[string]string,  map[string]string) {
	wins := map[string]string{
		"Rock": "Scissors",
		"Paper": "Rock",
		"Scissors": "Paper",
	}
	losses := map[string]string{}
	for k, v := range wins {
		losses[v] = k
	}
	return wins, losses
}

func determinePlay(opponent string, outcome int) int {
	wins, losses := getWinsAndLosses()
	if outcome == 3 {
		return getPlayScore(opponent) + outcome
	}
	if outcome == 0 {
		return getPlayScore(wins[opponent]) + outcome
	}
	if outcome == 6 {
		return getPlayScore(losses[opponent]) + outcome
	}
	panic(fmt.Sprintf("Invalid round scenario opponent: %s, outcome: %d", opponent, outcome))
}

func playGame(opponent, you string) int {
	if opponent == you {
		return 3
	}
	wins, losses := getWinsAndLosses()
	if wins[opponent] == you {
		return 6
	}
	if losses[opponent] == you {
		return 0
	}
	panic(fmt.Sprintf("Invalid move played %s vs %s", opponent, you))
}

func getPlayScore(play string) int {
	playScores := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}
	score, ok := playScores[play]
	if !ok {
		panic(fmt.Sprintf("%s is not a valid play", play))
	}
	return score
}

func playRound(opponent, you string) int {
	return playGame(opponent, you) + getPlayScore(you)
}

func GetResult1(rounds []string) int {
	result := 0
	for _, round := range rounds {
		roundMoves := strings.Split(round, " ")
		roundResult := playRound(decodePlay(roundMoves[0]), decodePlay(roundMoves[1]))
		result = result + roundResult
	}
	return result
}

func GetResult2(rounds []string) int {
	result := 0
	for _, round := range rounds {
		roundMoveOutcome := strings.Split(round, " ")
		opponentPlay := decodePlay(roundMoveOutcome[0])
		outcome := decodeOutcome(roundMoveOutcome[1])
		roundResult := determinePlay(opponentPlay, outcome)
		result = result + roundResult
	}
	return result
}

func Run() {
	input := utils.ReadInput(2)
	rounds := strings.Split(input, "\n")
	fmt.Printf("Day 02 part 1 result is:\n%d\n", GetResult1(rounds))
	fmt.Printf("\nDay 02 part 2 result is:\n%d\n", GetResult2(rounds))
}
