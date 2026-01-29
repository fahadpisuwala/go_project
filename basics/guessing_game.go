package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1
	println("Welcome to the Guessing Game!")
	println("I have selected a number between 1 and 100. Can you guess it?")
	println("The target number is:", source)
	println("The target number is:", random) // For testing purposes; remove in production

	println("The target number is:", target) // For testing purposes; remove in production
	// For testing purposes; remove in production
	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scan(&guess)

		if guess == target {
			println("Congratulations! You guessed the number correctly.")
			break
		} else if guess > target {
			println("Too high! Try a lower number.")
		} else {
			println("Too low! Try a higher number.")
		}

	}
}
