package main

import (
	"fmt"
	"rock-paper-scissors/utils"

	"github.com/fatih/color"
)

func main() {
	var totalScore = map[string]int{"Computer": 0, "User": 0, "Tie": 0, "Rounds": 0}
	var playAgain bool = true

	for playAgain {
		userChoice := utils.GetUserChoice()
		computerChoice := utils.GenerateComputerChoice()

		fmt.Println("User choice: ", userChoice)
		fmt.Println("Computer choice: ", computerChoice)

		winner := utils.CalculateWinner(computerChoice, userChoice)
		totalScore[winner]++
		totalScore["Rounds"]++

		playAgain = utils.PlayAgain()
	}

	fmt.Println("Final results")
	fmt.Println("--------------------------------------------")
	fmt.Printf("User: %v\n", totalScore["User"])
	fmt.Printf("Computer: %v\n", totalScore["Computer"])
	fmt.Printf("Tie: %v\n", totalScore["Tie"])
	fmt.Printf("Number of games played: %v\n", totalScore["Rounds"])

	switch {

	case totalScore["User"] > totalScore["Computer"]:
		color.Green("Final result: Congratulations, you won overall!")

	case totalScore["User"] < totalScore["Computer"]:
		color.Red("Final result: The computer did better overall. Better luck next time.")

	case totalScore["User"] == totalScore["Computer"]:
		color.Yellow("Final result: It was a tie overall.")

	default:
		color.Green("Thanks for playing!")
	}
}
