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

// start of the game. Encoding the secret word.
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

// guessing until user reveales the word or loses.
func makeGuess(hidWord, word string) {
	time.Sleep(time.Second)
	printCurrent(hidWord)

	// asks for the letter to guess.
	guess := askLetter(word)

	// user chose to guess the whole word.
	if guess == "GUESS" {
		guessWord(guess, word)
		makeGuess(hidWord, word)
	}

	// chosen letter is incorrect. Substracting 1 from lives.
	if !strings.Contains(word, guess) {
		decrementLives(word)
		makeGuess(hidWord, word)
	}

	// chosen letter is already open. Doesn't affect the lives.
	if strings.Contains(hidWord, guess) {
		fmt.Println("This letter is already open!")
		makeGuess(hidWord, word)
	}

	// opens the guessed letter in hidden word.
	hidWord = openLetters(hidWord, word, guess)

	// if no hidden letters left -> win. 
	if !strings.Contains(hidWord, "_") {
		fmt.Println(color.Ize(color.Green, "Word revealed!!"))
		fmt.Println("---------", word, "---------")
		finishGame(1)
	}

	// nolint: go-staticcheck
	makeGuess(hidWord, word) // has to be infinite.
}

// prints current lives count and an image.
func printCurrent(hiddenWord string) {
	fmt.Println()

	loadAnimation.Start()
	time.Sleep(2 * time.Second)
	loadAnimation.Stop()

	fmt.Printf("Currently you have %v lives left.\n", lives)
	fmt.Println(color.Ize(color.Yellow, visualLives[lives]))
	fmt.Println(color.Ize(color.Cyan, hiddenWord))
	fmt.Println()
}

// asks for a letter (or key word "guess"/"exit") from a user to guess.
func askLetter(word string) string {
	fmt.Println("Guess a letter: ")
	var letter string
	fmt.Scanln(&letter)

	letter = strings.ToUpper(letter)
	// fmt.Println("Your guess: ", letter) // usage: for the user to see if input was taken correctly.

	// key word to exit the game.
	if letter == "EXIT" {
		exitGame(1)
	}

	// key word to guess the whole word.
	if letter == "GUESS" {
		return letter
	}

	// not 1 character was typed; character isn't valid (only english leters are).
	if len(letter) > 1 || len(letter) < 1 || !strings.Contains(alphabet, letter) {
		fmt.Println(color.Ize(color.Purple, "Oh... Looks like you've entered an invalid letter."))
		fmt.Println("Please, try again.")
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

	fmt.Println("You've chosen to guess the word right away. Brave enough!")
	fmt.Println(color.Ize(color.Green, "Please, enter your guess:"))

	fmt.Scanln(&guessWord)

	// word is correct.
	if strings.ToUpper(guessWord) == word {
		finishGame(1) // 1 as code for a win.
	}

	decrementLives(word)
}

// taking away a life since the guess was valid, but incorrect.
func decrementLives(word string) {
	fmt.Println(color.Ize(color.Red, "Oh-oh. Incorrect. I'm taking away one of your lives!"))
	lives--

	// if no lives left - game lost.
	if lives < 1 {
		fmt.Println("Oh... You lost!")
		fmt.Println("The word was ", color.Ize(color.Cyan, word))
		finishGame(0) // 0 as code for a lose.
	}
}

// finishes the game with corresponding (win/lost) output.
func finishGame(res int) {
	switch res {
	case 0:
		fmt.Println()
		fmt.Println("You were so close!")

		askToPlayAgain()

	case 1:
		fmt.Println()
		fmt.Println(color.Ize(color.Green, "YAY!! You won!"))
		fmt.Println("Congratulations!")

		askToPlayAgain()

	default:
		log.Fatal("an undefined error occured while finishing the game...")
		
	}
}

// opens guessed letters in hidden word.
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