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
	fmt.Println("Вітаємо у Шибениці!")
	printMinPlusLine()
}

// tells the user game rules.
func printRules() {
	fmt.Println()
	fmt.Println(color.Ize(color.Cyan, "Оскільки ви бажаєте зіграти, я нагадаю вам деякі правила:"))
	printMinLine()
	time.Sleep(2 * time.Second)

	fmt.Println("Я оберу випадкове слово і ви матимете вгадати його літера за літерою.")
	time.Sleep(time.Second)
	fmt.Println("На початку кожної нової гри ви матимете 7 життів. Кожного разу, як ви робитимете помилку, я забиратиму 1!")
	time.Sleep(time.Second)
	fmt.Println("Система приймає українські літери як у нижньому, так і у ВЕРХНЬОМУ регістрі :)")
	time.Sleep(time.Second)
	fmt.Println("Якщо ви захочете вгадати ціяде слово одразу, ви можете це зробити, попередньо ввівши команду")
	time.Sleep(time.Second)
	fmt.Println(`"вгадати" замість літери, а потім - саме слово.`)
	time.Sleep(time.Second)
	fmt.Println(`Ви завжди можете вийти з гри, ввівши команду "вийти" (прогрес гри не зберігатиметься).`)
	time.Sleep(time.Second)
	printMinLine()

	fmt.Println("На цьому все. Удачі!")
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

	fmt.Println(color.Ize(color.Purple, "Дайте-но мені вигадати якесь цікаве слово..."))
}

// reveals the word since all letters are now open.
func wordRevealed(word string) {
	fmt.Println(color.Ize(color.Green, "Слово розкрито!!"))
	fmt.Println("---------", color.Ize(color.Cyan, word), "---------")
}

// prints current lives count and an image.
func printCurrent(hiddenWord string) {
	fmt.Println()

	loadAnimation.Start()
	time.Sleep(2 * time.Second)
	loadAnimation.Stop()

	fmt.Printf("Наразі ви маєте ще %v життів.\n", lives)
	fmt.Println(color.Ize(color.Yellow, visualLives[lives]))
	fmt.Println(color.Ize(color.Cyan, hiddenWord))
	fmt.Println()
}

// tells the user if they entered an invalid character.
func invalidLetter() {
	fmt.Println(color.Ize(color.Purple, "Ой... Ви обрали неприпустиму літеру."))
	fmt.Println("Будь ласка, спробуйте ще раз.")
}

// offers to guess the whole word.
func offerToGuess() {
	fmt.Println("Ви обрали вгадати ціле слово одразу. Сміливо!")
	fmt.Println(color.Ize(color.Green, "Будь ласка, введіть ваше припущення:"))
}

// tells the user they lost and prints out the word.
func printLoss(word string) {
	fmt.Println("О ні... Ви програли!")
	fmt.Println("Загадане слово: ", color.Ize(color.Cyan, word))
	fmt.Println()
	fmt.Println("Ви були близько!")
}

// tells the user they won.
func printWin() {
	fmt.Println()
	fmt.Println(color.Ize(color.Green, "ВАУ! Ви перемогли!"))
	fmt.Println("Вітаю!")
}

// says goodbye to the user: sad if no game played, otherwise - happy.
func sadBye() {
	fmt.Println("OK. Сподіваюсь, що ви ще зіграєте...")
	time.Sleep(time.Second)
	fmt.Println(color.Ize(color.Blue, "Приходьте ще, будь ласка..."))
}

func happyBye() {
	fmt.Println("OK. Сподіваюсь ви отримали задоволення від проведеного разом часу!")
	time.Sleep(time.Second)
	fmt.Println(color.Ize(color.Purple, "Приходьте ще ~"))
}

// tells the user the option they've chosen is invalid and asks to trye again.
func invalidOption() {
	fmt.Println(color.Ize(color.Red, "Ой... Ви обрали варіант, що не підтримується системою!"))
	fmt.Println(color.Ize(color.Yellow, "==== failed to start the game: invalid option ===="))
	time.Sleep(time.Second)
	fmt.Println("Будь ласка, спробуйте ще раз...")
}