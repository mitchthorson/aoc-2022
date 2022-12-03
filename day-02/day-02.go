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

func playGame(opponent, you string) int {
	if opponent == you {
		return 3
	}
	if opponent == "Rock" {
		if you == "Paper" {
			return 6
		}
		if you == "Scissors" {
			return 0
		}
	}
	if opponent == "Paper" {
		if you == "Rock" {
			return 0
		}
		if you == "Scissors" {
			return 6
		}
	}
	if opponent == "Scissors" {
		if you == "Rock" {
			return 6
		}
		if you == "Paper" {
			return 0
		}
	}
	panic(fmt.Sprintf("Invalid move played %s vs %s", opponent, you))
}

func playRound(opponent, you string) int {
	playScores := map[string] int{
		"Rock": 1,
		"Paper": 2,
		"Scissors": 3,
	}
	return playGame(opponent, you) + playScores[you]
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

func Run() {
	input := utils.ReadInput(2)
	rounds := strings.Split(input, "\n")
	fmt.Printf("\nDay 02 part 1 result is:\n%d\n", GetResult1(rounds))
}
