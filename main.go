package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	playerScore, computerScore := 0, 0

	clearScreen(os.Stdout)

	fmt.Println(`Welcome to "Rock, Paper & Scissors"`)
	fmt.Println(`Game is played before anybody reaches 3 scores.`)
	fmt.Println(`Print 'r(ock)', 'p(aper)' or 's(cissors)' for rock, paper or scissors.`)
	fmt.Print("----------\n")

	for playerScore < 3 && computerScore < 3 {
		rand.Seed(time.Now().UnixNano())
		fmt.Print("\nYour turn: ")
		playerChoice, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
		}

		playerChoice = playerChoice[:len(playerChoice)-1]
		playerChoiceNum, err := convertPlayerChoiceToNum(playerChoice)

		if err != nil {
			fmt.Println(err)
			continue
		}

		computerChoiceNum := rand.Int() % 3
		computerChoice := covertComputerChoiceToString(computerChoiceNum)
		fmt.Printf("Computer chooses: %s\n", computerChoice)

		currPlayerScore, currComputerScore := roundWinner(playerChoiceNum, computerChoiceNum)
		playerScore += currPlayerScore
		computerScore += currComputerScore
		fmt.Printf("Score: %d:%d\n", playerScore, computerScore)
	}

	totalResult(playerScore, computerScore)
}

func clearScreen(w io.Writer) {
	cmd := exec.Command(`clear`)
	cmd.Stdout = w
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func roundWinner(plChoice, compChoice int) (int, int) {
	if plChoice == compChoice {
		return 0, 0
	} else if (plChoice+1)%3 == compChoice {
		return 0, 1
	} else {
		return 1, 0
	}
}

func convertPlayerChoiceToNum(playerChoice string) (int, error) {
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

func covertComputerChoiceToString(computerChoiceNum int) string {
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

func totalResult(playerScore, computerScore int) {
	fTotalResult(os.Stdout, playerScore, computerScore)
}

func fTotalResult(w io.Writer, playerScore, computerScore int) {
	fmt.Fprintln(w, "\n----------")
	if playerScore > computerScore {
		fmt.Fprintln(w, "Human wins!")
	} else if playerScore == computerScore {
		fmt.Fprintln(w, "Draw.")
	} else {
		fmt.Fprintln(w, "Computer wins!")
	}
}