package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

// ask to play for the first time.
func askToPlay() {
	loadAnimation.Stop()

	fmt.Println("Would you like to play? (y/n)")
	var ans string
	fmt.Scanln(&ans)

	loadAnimation.Start()
	time.Sleep(2 * time.Second)

	switch strings.ToLower(ans) {
	case "y": 
		mainMenu() // start the game from the main menu.

	case "n":
		exitGame(0) // exits the game: 0 as code for "haven't played".

	default:
		// re-asks since input was invalid.
		loadAnimation.Stop()

		fmt.Println(color.Ize(color.Red, "Oh-oh... It looks like you've chosen an invalid option!"))
		fmt.Println(color.Ize(color.Yellow, "==== failed to start the game: invalid option ===="))
		time.Sleep(time.Second)
		fmt.Println("Please, try again...")

		askToPlay()
	}
}

// asks to play. Repeating code since different output is needed.
func askToPlayAgain() {
	fmt.Println("Would you like to play again? (y/n)")
	var ans string
	fmt.Scanln(&ans)

	loadAnimation.Start()
	time.Sleep(3 * time.Second)

	switch strings.ToLower(ans) {
	case "y":
		startTheGame() // starts the game directly.

	case "n":
		exitGame(1) //exits the game; 1 as code for "played some games".

	default:
		// re-asks to play since the input was invalid.
		loadAnimation.Stop()

		fmt.Println(color.Ize(color.Red, "Oh-oh... It looks like you've chosen an invalid option!"))
		fmt.Println(color.Ize(color.Yellow, "==== failed to start the game: invalid option ===="))
		time.Sleep(time.Second)
		fmt.Println("Please, try again...")

		askToPlayAgain()
	}
}

// exits the game with corresponding output. 
func exitGame(played int) {
	loadAnimation.Stop()

	fmt.Println()
	fmt.Println(color.Ize(color.Yellow, "-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-"))

	switch played {
	case 0: // no games played
		fmt.Println("OK. Hope you'll come back and play for once...")
		time.Sleep(time.Second)
		fmt.Println(color.Ize(color.Blue, "See you soon..."))
		
	case 1: // some games played
		fmt.Println("OK. Hope you've enjoyed the time spent!")
		time.Sleep(time.Second)
		fmt.Println(color.Ize(color.Purple, "See you soon ~"))

	default:
		log.Fatal("an undefined error occured while exiting the game")
	}

	fmt.Println(color.Ize(color.Yellow, "-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-"))
	fmt.Println()	

	// stops program execution.
	os.Exit(0)
}
