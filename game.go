package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var visualLives = map[int]string{
	0: "",
	1: `
	|
	|
	|
	|
	|

	`,
	2: `
	___________
	| /        
	|/        
	|          
	|          
	|

	`,
	3: `
	___________
	| /        | 
	|/        
	|          
	|          
	|

	`,
	4: `
	___________
	| /        | 
	|/        ( )
	|          
	|          
	|
	
	`,
	5: `
	___________
	| /        | 
	|/        ( )
	|          |
	|          
	|
	
	`,
	6: `
	___________
	| /        | 
	|/        ( )
	|          |
	|         / 
	|
	
	`,
	7: `
	___________
	| /        | 
	|/        ( )
	|          |
	|         / \\
	|
	
	`,
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lives = 7

func startGuessing(word string) int {
	lives = 7

	var hidWord string
	for i := 0; i < len(word); i++ {
		hidWord += "_ "
	}

	word = strings.ToUpper(word)
	makeGuess(hidWord, word)
	
	return 0
}

func makeGuess(hidWord, word string) {
	time.Sleep(time.Second)
	printCurrent(hidWord)

	guess := askLetter(word)

	if guess == "guess" {
		guessWord(guess, word)
		makeGuess(hidWord, word)
	}

	if !strings.Contains(word, guess) {
		decrementLives()
		makeGuess(hidWord, word)
	}

	if strings.Contains(hidWord, guess) {
		fmt.Println("This letter is already open!")
		makeGuess(hidWord, word)
	}

	hidWord = openLetters(hidWord, word, guess)

	if !strings.Contains(hidWord, "_") {
		fmt.Println("Word revealed!!")
		fmt.Println(word)
		finishGame(1)
	}

	// nolint: go-staticcheck
	makeGuess(hidWord, word)
}

func printCurrent(hiddenWord string) {
	fmt.Printf("Currently you have %v lives left\n", lives)
	fmt.Println(visualLives[lives])
	fmt.Println(hiddenWord)
	fmt.Println()
}

func askLetter(word string) string {
	fmt.Println("Guess a letter: ")
	var letter string
	fmt.Scanln(&letter)

	letter = strings.ToUpper(letter)

	if letter == "EXIT" {
		exitGame(1)
	}

	if letter == "GUESS" {
		return letter
	}

	if len(letter) > 1 || len(letter) < 1 || !strings.Contains(alphabet, letter) {
		fmt.Println("Oh... Looks like you've entered an invalid letter.")
		fmt.Println("Please, try again.")
		askLetter(word)
	}
	
	return letter
}

func guessWord(guess, word string) {
	if guess == word {
		finishGame(1)
	}
	decrementLives()
}

func decrementLives() {
	fmt.Println("Oh-oh. Incorrect. I'm taking away one of your lives!")
	lives--

	if lives < 1 {
		finishGame(0)
	}
}

func finishGame(res int) {
	switch res {
	case 0:
		fmt.Println(visualLives[lives])
		fmt.Println("What a pity! You lost... You were so close!")
		askToPlayAgain()
	case 1:
		fmt.Println("YAY!! You won!")
		fmt.Println("Congratulations!")
		askToPlayAgain()
	default:
		log.Fatal("an undefined error occures while finishing the game...")
		
	}
}

func openLetters(hidWord string, word, guess string) string {
	var wordIndexes []int
	for i := 0; i < len(word); i++ {
		if string(word[i]) == guess {
			wordIndexes = append(wordIndexes, i)
		}
	}

	for _, ind := range wordIndexes {
		hidWord = hidWord[:2*ind] + guess + hidWord[2*ind + 1:]
	}

	return hidWord
}