package main

import (
	"bufio"
	"fmt"
	rps "github.com/LightBulbfromSpace/rock-paper-scissors-console-game/pkg"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	playerScore, computerScore := 0, 0

	clearScreen(os.Stdout)

	fmt.Println(`Welcome to "Rock, Paper & Scissors"`)
	fmt.Println(`Game is played before anybody reaches 3 scores.`)
	fmt.Println(`Print 'r(ock)', 'p(aper)' or 's(cissors)' for rock, paper or scissors.`)
	fmt.Print("----------\n")

	continueGame := true

	for continueGame {
		for playerScore < 3 && computerScore < 3 {
			rand.Seed(time.Now().UnixNano())
			fmt.Print("\nYour turn: ")
			playerChoice, err := reader.ReadString('\n')

			if err != nil {
				log.Println(err)
			}

			playerChoice = playerChoice[:len(playerChoice)-1]
			playerChoiceNum, err := rps.ConvertPlayerChoiceToNum(playerChoice)

			if err != nil {
				fmt.Println(err)
				continue
			}

			computerChoiceNum := rand.Int() % 3
			computerChoice := rps.CovertComputerChoiceToString(computerChoiceNum)
			fmt.Printf("Computer chooses: %s\n", computerChoice)

			currPlayerScore, currComputerScore := rps.RoundWinner(playerChoiceNum, computerChoiceNum)
			playerScore += currPlayerScore
			computerScore += currComputerScore
			fmt.Printf("Score: %d:%d\n", playerScore, computerScore)
		}

		rps.TotalResult(playerScore, computerScore)

		fmt.Println("Do you want to play again? Press [y/any character]")
		continueGameStr, _ := reader.ReadString('\n')
		if continueGameStr == "y\n" {
			continueGame = true
			playerScore, computerScore = 0, 0
		} else {
			continueGame = false
		}
	}
}

func clearScreen(w io.Writer) {
	var cmd *exec.Cmd
	os := runtime.GOOS

	switch os {
	case "linux":
		cmd = exec.Command(`clear`)
	case "windows":
		cmd = exec.Command(`cls`)
	default:
		return
	}

	cmd.Stdout = w
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
