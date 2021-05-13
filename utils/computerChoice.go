package utils

import (
	"math/rand"
	"time"
)

func GenerateComputerChoice() (choice string) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)

	switch randNum {
	case 0:
		choice = "rock"

	case 1:
		choice = "paper"

	case 2:
		choice = "scissors"
	}

	return
}
