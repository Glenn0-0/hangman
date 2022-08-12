package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	loadAnimation := spinner.New(spinner.Charset[10], 100*time.Millisecond)
	loadAnimation.Start()
	time.Sleep(2 * time.Second) 
	fmt.Println("Welcome to Hangman Game!")
	loadAnimation.Stop()
}