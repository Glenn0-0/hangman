package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func askToPlay() {
	loadAnimation.Stop()

	fmt.Println("Would you like to play? (y/n)")
	var ans string
	fmt.Scanln(&ans)

	loadAnimation.Start()
	time.Sleep(2 * time.Second)

	switch ans {
	case "y": 
		mainMenu()

	case "n":
		exitGame(0)

	default:
		loadAnimation.Stop()

		fmt.Println("Oh-oh... It looks like you've chosen an invalid option!")
		fmt.Println("==== failed to start the game: invalid option ====")
		time.Sleep(time.Second)
		fmt.Println("Please, try again...")

		askToPlay()
	}
}

func askToPlayAgain() {
	fmt.Println("Would you like to play again? (y/n)")
	var ans string
	fmt.Scanln(&ans)

	loadAnimation.Start()
	time.Sleep(3 * time.Second)

	switch ans {
	case "y":
		startTheGame()

	case "n":
		exitGame(1)

	default:
		loadAnimation.Stop()

		fmt.Println("Oh-oh... It looks like you've chosen an invalid option!")
		fmt.Println("==== failed to start the game: invalid option ====")
		time.Sleep(time.Second)
		fmt.Println("Please, try again...")

		askToPlayAgain()
	}
}

func exitGame(played int) {
	loadAnimation.Stop()

	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")

	switch played {
	case 0:
		fmt.Println("OK. Hope you'll come back and play for once...")
		time.Sleep(time.Second)
		fmt.Println("See you soon...")
		
	case 1:
		fmt.Println("OK. Hope you've enjoyed the time spent!")
		time.Sleep(time.Second)
		fmt.Println("See you soon ~")

	default:
		log.Fatal("an undefined error occured while exiting the game")
	}

	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	fmt.Println()	

	os.Exit(0)
}
