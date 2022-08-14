package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/TwiN/go-color"
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
	|/        ( ) <-- путін
	|          
	|          
	|
	
	`,
	5: `
	___________
	| /        | 
	|/        ( ) <-- путін
	|          |
	|          
	|
	
	`,
	6: `
	___________
	| /        | 
	|/        ( ) <-- путін
	|          |
	|         / 
	|
	
	`,
	7: `
	___________
	| /        | 
	|/        ( ) <-- путін
	|          |
	|         / \\
	|
	
	`,
}

const alphabet = "АБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ"
var lives = 7

// start of the game. Encoding the secret word.
func startGuessing(word string) int {
	lives = 7

	var hidWord string
	for i := 0; i < len(word); i += 2 {
		hidWord += "__  "
	}

	word = strings.ToUpper(word)
	makeGuess(hidWord, word)
	
	return 0
}

// guessing until user reveals the word or loses.
func makeGuess(hidWord, word string) {
	time.Sleep(time.Second)
	printCurrent(hidWord)

	// asks for the letter to guess.
	guess := askLetter(word)

	// user chose to guess the whole word.
	if guess == "ВГАДАТИ" {
		guessWord(guess, word)
		makeGuess(hidWord, word)
	}

	// chosen letter is incorrect. Subtracting 1 from lives.
	if !strings.Contains(word, guess) {
		decrementLives(word)
		makeGuess(hidWord, word)
	}

	// chosen letter is already open. Doesn't affect the lives.
	if strings.Contains(hidWord, guess) {
		fmt.Println("Ця літера вже відкрита!")
		makeGuess(hidWord, word)
	}

	// opens the guessed letter in hidden word.
	hidWord = openLetters(hidWord, word, guess)

	// if no hidden letters left -> win. 
	if !strings.Contains(hidWord, "_") {
		wordRevealed(word)
		finishGame(1)
	}

	// nolint: go-staticcheck
	makeGuess(hidWord, word) // has to be infinite.
}

// asks for a letter (or key word "guess"/"exit") from a user to guess.
func askLetter(word string) string {
	fmt.Println("Введіть літеру: ")
	var letter string
	fmt.Scanln(&letter)

	letter = strings.ToUpper(letter)
	// fmt.Println("Your guess: ", letter) // usage: for the user to see if input was taken correctly.

	// key word to exit the game.
	if letter == "ВИЙТИ" {
		exitGame(1)
	}

	// key word to guess the whole word.
	if letter == "ВГАДАТИ" {
		return letter
	}

	// not 1 character was typed; character isn't valid (only english letters are).
	if len(letter) > 2 || len(letter) < 2 || !strings.Contains(alphabet, letter) {
		invalidLetter()
		askLetter(word)
	}
	
	return letter
}

// guessing the whole word at once.
func guessWord(guess, word string) {
	var guessWord string
	fmt.Println()

	// decorative animation.
	loadAnimation.Start()
	time.Sleep(2 * time.Second)
	loadAnimation.Stop()

	offerToGuess()

	fmt.Scanln(&guessWord)

	// word is correct.
	if strings.ToUpper(guessWord) == word {
		finishGame(1) // 1 as code for a win.
	}

	decrementLives(word)
}

// taking away a life since the guess was valid, but incorrect.
func decrementLives(word string) {
	fmt.Println(color.Ize(color.Red, "Ой... На жаль, невірно. Я забираю у вас одне життя!"))
	lives--

	// if no lives left - game lost.
	if lives < 1 {
		printLoss(word)
		finishGame(0) // 0 as code for a lose.
	}
}

// finishes the game with corresponding (win/lost) output.
func finishGame(res int) {
	switch res {
	case 0:
		askToPlayAgain()

	case 1:
		printWin()
		askToPlayAgain()

	default:
		log.Fatal("an undefined error occurred while finishing the game...")
		
	}
}

// opens guessed letters in hidden word.
func openLetters(hidWord, word, guess string) string {
	indexes := []int{}
	numOfLetters := strings.Count(word, guess)

	for i := 0; i < numOfLetters; i++ {
		index := strings.Index(word, guess)
		indexes = append(indexes, index)
		word = strings.Replace(word, guess, "Ы", 1)
	}

	for _, index := range indexes {
		hidWord = hidWord[:2 * index] + guess + hidWord[2 * index + 2:]
	}

	return hidWord
}