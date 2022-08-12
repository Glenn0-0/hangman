package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var loadAnimation = spinner.New(spinner.CharSets[27], 100*time.Millisecond)
const filePath = "./lib/words.txt"

func main() {
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	fmt.Println("Welcome to Hangman Game!")
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")

	loadAnimation.Start()
	time.Sleep(3 * time.Second)
	askToPlay()
}

func mainMenu() {
	loadAnimation.Stop()

	fmt.Println("Since you want to play, I'll remind you of some rules:")
	fmt.Println()
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
	fmt.Println(`You can always exit the game (progress isn't being saved) by entering "exit"`)
	time.Sleep(time.Second)

	fmt.Println("That's it. Good luck!")

	loadAnimation.Start()
	time.Sleep(4 * time.Second)
	startTheGame()
}

func startTheGame() {
	randWord := chooseSecretWord()

	gameResult := startGuessing(randWord)
	finishGame(gameResult)
}

func chooseSecretWord() string {
	loadAnimation.Stop()
	wordList := readLines()

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

	fmt.Println("Let me think of an interesting word...")
	loadAnimation.Start()
	time.Sleep(3 * time.Second)
	
	rand.Seed(time.Now().Unix())
	randWord := wordList[rand.Intn(len(wordList) - 1)]

	loadAnimation.Stop()
	
	return randWord
}

func readLines() []string {
	var wordList []string

	readFile, errRead := os.Open(filePath)
	if errRead != nil {
		log.Fatal(errRead)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		wordList = append(wordList, fileScanner.Text())
	} 

	return wordList
}