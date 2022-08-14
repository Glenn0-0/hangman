package main

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
)

// decorations.
func printMinPlusLine() {
	fmt.Println(color.Ize(color.Yellow, "-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-"))
}

func printMinLine() {
	fmt.Println(color.Ize(color.Gray, "-------------------------------------------------------------------"))
}

// greets the user.
func greet() {
	printMinPlusLine()
	fmt.Println("Welcome to Hangman Game!")
	printMinPlusLine()
}

// tells the user game rules.
func printRules() {
	fmt.Println()
	fmt.Println(color.Ize(color.Cyan, "Since you want to play, I'll remind you of some rules:"))
	printMinLine()
	time.Sleep(2 * time.Second)

	fmt.Println("I'll randomly choose a word, and you'll have to guess it letter by letter.")
	time.Sleep(time.Second)
	fmt.Println("You have 7 lives. Each time you enter an incorrect letter or word, I'll take 1 life from you!")
	time.Sleep(time.Second)
	fmt.Println("You can enter letters in both lower and upper case.")
	time.Sleep(time.Second)
	fmt.Println("If you will have an early guess, which word I've chosen, you can check if it's correct")
	time.Sleep(time.Second)
	fmt.Println(`by entering "guess" instead of a letter.`)
	time.Sleep(time.Second)
	fmt.Println(`You can always exit the game (progress isn't being saved) by entering "exit".`)
	time.Sleep(time.Second)

	fmt.Println("That's it. Good luck!")
}

// imitates thinking process while choosing a random word.
func animateThinking() {
	fmt.Println()
	fmt.Println("*")
	time.Sleep(time.Second)
	fmt.Println("*")
	time.Sleep(time.Second)
	fmt.Println("*")
	time.Sleep(time.Second)
	fmt.Println("*")
	time.Sleep(time.Second)
	fmt.Println("*")
	time.Sleep(time.Second)

	fmt.Println(color.Ize(color.Purple, "Let me think of an interesting word..."))
}

// reveals the word since all letters are now open.
func wordRevealed(word string) {
	fmt.Println(color.Ize(color.Green, "Word revealed!!"))
	fmt.Println("---------", color.Ize(color.Cyan, word), "---------")
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

// tells the user if they entered an invalid character.
func invalidLetter() {
	fmt.Println(color.Ize(color.Purple, "Oh... Looks like you've entered an invalid letter."))
		fmt.Println("Please, try again.")
}

// offers to guess the whole word.
func offerToGuess() {
	fmt.Println("You've chosen to guess the word right away. Brave enough!")
	fmt.Println(color.Ize(color.Green, "Please, enter your guess:"))
}

// tells the user they lost and prints out the word.
func printLoss(word string) {
	fmt.Println("Oh... You lost!")
	fmt.Println("The word was ", color.Ize(color.Cyan, word))
	fmt.Println()
	fmt.Println("What a pity... You were so close!")
}

// tells the user they won.
func printWin() {
	fmt.Println()
	fmt.Println(color.Ize(color.Green, "YAY!! You won!"))
	fmt.Println("Congratulations!")
}

// says goodbye to the user: sad if no game played, otherwise - happy.
func sadBye() {
	fmt.Println("OK. Hope you'll come back and play for once...")
	time.Sleep(time.Second)
	fmt.Println(color.Ize(color.Blue, "See you soon..."))
}

func happyBye() {
	fmt.Println("OK. Hope you've enjoyed the time spent!")
	time.Sleep(time.Second)
	fmt.Println(color.Ize(color.Purple, "See you soon ~"))
}

// tells the user the option they've chosen is invalid and asks to trye again.
func invalidOption() {
	fmt.Println(color.Ize(color.Red, "Oh-oh... It looks like you've chosen an invalid option!"))
	fmt.Println(color.Ize(color.Yellow, "==== failed to start the game: invalid option ===="))
	time.Sleep(time.Second)
	fmt.Println("Please, try again...")
}