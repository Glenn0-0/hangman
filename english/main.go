package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var loadAnimation = spinner.New(spinner.CharSets[36], 100*time.Millisecond)
const filePath = "./lib/words.txt"

func main() {
	// greetings.
	greet()

	// decorative animation.
	loadAnimation.Start()
	time.Sleep(3 * time.Second)

	askToPlay()
}

func mainMenu() {
	loadAnimation.Stop()

	// game rules.
	printRules()

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

	animateThinking()

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