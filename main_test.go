package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type Result struct {
	playerScore   int
	computerScore int
}

func TestRoundWinner(t *testing.T) {
	cases := []struct {
		playerChoice   int
		computerChoice int
		want           Result
	}{
		{
			ROCK,
			SCISSORS,
			Result{1, 0},
		},
		{
			SCISSORS,
			PAPER,
			Result{1, 0},
		},
		{
			SCISSORS,
			SCISSORS,
			Result{0, 0},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			plRes, compRes := roundWinner(tc.playerChoice, tc.computerChoice)
			if plRes != tc.want.playerScore && compRes != tc.want.computerScore {
				t.Errorf("got player: %d, computer: %d, but want %v", plRes, compRes, tc.want)
			}
		})
	}
}

func TestTotalResult(t *testing.T) {
	cases := []struct {
		totalPlSc   int
		totalCompSc int
		expected    string
	}{
		{3, 1, "\n----------\nHuman wins!\n"},
		{2, 3, "\n----------\nComputer wins!\n"},
		{1, 1, "\n----------\nDraw.\n"},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			buffer := bytes.Buffer{}
			fTotalResult(&buffer, tc.totalPlSc, tc.totalCompSc)
			if !reflect.DeepEqual(buffer.String(), tc.expected) {
				t.Errorf("expected %s, but got %s", tc.expected, buffer.String())
			}
		})
	}
}

func TestCovertComputerChoiceToString(t *testing.T) {
	cases := []struct {
		compChoiceNum int
		expected      string
	}{
		{ROCK, "ROCK"},
		{PAPER, "PAPER"},
		{SCISSORS, "SCISSORS"},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			got := covertComputerChoiceToString(tc.compChoiceNum)
			if got != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, got)
			}
		})
	}
}

func TestConvertPlayerChoiceToNum(t *testing.T) {
	t.Run("success conversion", func(t *testing.T) {
		cases := []struct {
			playerChoice string
			expected     int
		}{
			{"r", ROCK},
			{"p", PAPER},
			{"s", SCISSORS},
		}
		for i, tc := range cases {
			t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
				got, err := convertPlayerChoiceToNum(tc.playerChoice)
				if err != nil {
					t.Error("Got error but didn't expect one.")
				}
				if got != tc.expected {
					t.Errorf("for %s expected %d, but got %d", tc.playerChoice, tc.expected, got)
				}
			})
		}
	})
	t.Run("error check", func(t *testing.T) {
		_, err := convertPlayerChoiceToNum("anything")
		if err == nil {
			t.Error("Expected to get error, but didn't get one.")
		}
	})
}
