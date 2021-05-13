package utils

import "github.com/fatih/color"

func CalculateWinner(computerChoice, userChoice string) (winner string) {
	if userChoice == "rock" {
		switch computerChoice {
		case "rock":
			winner = "Tie"
			color.Yellow("Both picked rock, it's a tie!")

		case "paper":
			winner = "Computer"
			color.Red("Paper covers rock...computer wins!")

		case "scissors":
			winner = "User"
			color.Green("Rock crushes scissors...you win!")
		}
	}

	if userChoice == "paper" {
		switch computerChoice {
		case "rock":
			winner = "User"
			color.Green("Paper covers rock...you win!")

		case "paper":
			winner = "Tie"
			color.Yellow("Both picked paper, it's a tie!")

		case "scissors":
			winner = "Computer"
			color.Red("Scissors cuts paper...computer wins!")
		}
	}

	if userChoice == "scissors" {
		switch computerChoice {
		case "rock":
			winner = "Computer"
			color.Red("Rock crushes scissors...computer wins!")

		case "paper":
			winner = "User"
			color.Green("Scissors cuts paper...you win!")

		case "scissors":
			winner = "Tie"
			color.Yellow("Both picked scissors...it's a tie!")
		}
	}

	return
}
