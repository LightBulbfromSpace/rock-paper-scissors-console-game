package rps

import (
	"errors"
	"fmt"
	"io"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func RoundWinner(player1, player2 int) (resPlayer1 int, resPlayer2 int) {
	if player1 == player2 {
		return 0, 0
	} else if (player1+1)%3 == player2 {
		return 0, 1
	} else {
		return 1, 0
	}
}

func ConvertPlayerChoiceToNum(playerChoice string) (int, error) {
	var playerChoiceNum int
	switch playerChoice {
	case "r", "rock":
		playerChoiceNum = ROCK
		break
	case "p", "paper":
		playerChoiceNum = PAPER
		break
	case "s", "scissors":
		playerChoiceNum = SCISSORS
		break
	default:
		return -1, errors.New("wrong input")
	}
	return playerChoiceNum, nil
}

func CovertComputerChoiceToString(computerChoiceNum int) string {
	var computerChoice string
	switch computerChoiceNum {
	case 0:
		computerChoice = "ROCK"
		break
	case 1:
		computerChoice = "PAPER"
		break
	case 2:
		computerChoice = "SCISSORS"
		break
	}
	return computerChoice
}

func FTotalResult(w io.Writer, playerScore, computerScore int) {
	fmt.Fprintln(w, "\n----------")
	if playerScore > computerScore {
		fmt.Fprintln(w, "Human wins!")
	} else if playerScore == computerScore {
		fmt.Fprintln(w, "Draw.")
	} else {
		fmt.Fprintln(w, "Computer wins!")
	}
}
