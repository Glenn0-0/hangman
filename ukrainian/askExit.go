package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// ask to play for the first time.
func askToPlay() {
	loadAnimation.Stop()

	fmt.Println("Бажаєте зіграти? (так/ні)")
	var ans string
	fmt.Scanln(&ans)

	loadAnimation.Start()
	time.Sleep(2 * time.Second)

	switch strings.ToLower(ans) {
	case "так": 
		mainMenu() // start the game from the main menu.

	case "ні":
		exitGame(0) // exits the game: 0 as code for "haven't played".

	default:
		// re-asks since input was invalid.
		loadAnimation.Stop()
		invalidOption()
		askToPlay()
	}
}

// asks to play. Repeating code since different output is needed.
func askToPlayAgain() {
	fmt.Println("Бажаєте зіграти знову? (так/ні)")
	var ans string
	fmt.Scanln(&ans)

	loadAnimation.Start()
	time.Sleep(3 * time.Second)

	switch strings.ToLower(ans) {
	case "так":
		startTheGame() // starts the game directly.

	case "ні":
		exitGame(1) //exits the game; 1 as code for "played some games".

	default:
		// re-asks to play since the input was invalid.
		loadAnimation.Stop()
		invalidOption()
		askToPlayAgain()
	}
}

// exits the game with corresponding output. 
func exitGame(played int) {
	loadAnimation.Stop()

	fmt.Println()
	printMinPlusLine()

	switch played {
	case 0: // no games played
		sadBye()
		
	case 1: // some games played
		happyBye()

	default:
		log.Fatal("an undefined error occurred while exiting the game")
	}

	printMinPlusLine()
	fmt.Println()	

	// stops program execution.
	os.Exit(0)
}
