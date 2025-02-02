package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type GameConfig struct {
	MaxNumber     int
	Chances       int
	TargetNumber  int
	DifficultyMap map[string]int
}

func main() {
	gameConfig := &GameConfig{
		MaxNumber: 100,
		DifficultyMap: map[string]int{
			"easy":   10,
			"medium": 5,
			"hard":   3,
		},
	}

	fmt.Println("number-guess")
	fmt.Println(`Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have to guess the correct number.`)
	if err := runGame(gameConfig); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func runGame(config *GameConfig) error {
	difficulty := selectDifficulty()
	config.Chances = config.DifficultyMap[difficulty]
	config.TargetNumber = generateRandomNumber(config.MaxNumber)

	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty)
	fmt.Printf("You have %d chances to guess the correct number.\n", config.Chances)
	fmt.Println("Let's start the game!")

	for attempts := 1; attempts <= config.Chances; attempts++ {
		guess := getGuess(attempts)
		if guess == config.TargetNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			return nil
		}
		if guess < config.TargetNumber {
			fmt.Println("Incorrect! The number is greater than your guess.")
		} else {
			fmt.Println("Incorrect! The number is less than your guess.")
		}
		if attempts == config.Chances {
			fmt.Printf("Sorry, you've run out of chances. The correct number was %d.\n", config.TargetNumber)
		}
	}

	return nil
}

func selectDifficulty() string {
	fmt.Println("\nPlease select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	for {
		fmt.Print("\nEnter your choice: ")
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
		switch choice {
		case 1:
			return "easy"
		case 2:
			return "medium"
		case 3:
			return "hard"
		}
	}
}

func generateRandomNumber(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max) + 1
}

func getGuess(attempt int) int {
	for {
		fmt.Printf("Attempt %d - Enter your guess: ", attempt)
		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil || guess < 1 || guess > 100 {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
		}
		return guess
	}
}
