package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/TwiN/go-color"
)

var loadAnimation = spinner.New(spinner.CharSets[36], 100*time.Millisecond)
const filePath = "./lib/words.txt"

func main() {
	// greetings.
	fmt.Println(color.Ize(color.Yellow, "-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-"))
	fmt.Println("Welcome to Hangman Game!")
	fmt.Println(color.Ize(color.Yellow, "-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-"))

	// decorative animation.
	loadAnimation.Start()
	time.Sleep(3 * time.Second)

	askToPlay()
}

func mainMenu() {
	loadAnimation.Stop()

	// game rules.
	fmt.Println()
	fmt.Println(color.Ize(color.Cyan, "Since you want to play, I'll remind you of some rules:"))
	fmt.Println(color.Ize(color.Gray, "---------------------------------------------------------"))
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
	fmt.Println(color.Ize(color.Gray, "---------------------------------------------------------"))

	fmt.Println("That's it. Good luck!")

	// decorative animation.
	loadAnimation.Start()
	time.Sleep(4 * time.Second)

	startTheGame()
}

// starts the game if user chose "y".
func startTheGame() {
	// randomly choosing a word from library.
	randWord := chooseSecretWord()

	// getting 0 if lost, 1 if won. Passing to finish the game with corresponding output.
	gameResult := startGuessing(randWord)
	finishGame(gameResult)
}

// randomly chooses and returns a word from the library.
func chooseSecretWord() string {
	loadAnimation.Stop()
	wordList := readLines() // reads the file with words.

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
	loadAnimation.Start()
	time.Sleep(3 * time.Second)
	
	rand.Seed(time.Now().Unix())
	randWord := wordList[rand.Intn(len(wordList) - 1)]

	loadAnimation.Stop()
	
	return randWord
}

// opens, reads and returns all the words from library (filePath).
func readLines() []string {
	var wordList []string

	readFile, errRead := os.Open(filePath)
	if errRead != nil {
		log.Fatal(errRead)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		wordList = append(wordList, fileScanner.Text())
	} 

	return wordList
}