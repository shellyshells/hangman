package main

import (
	"bufio"
	"fmt"
	"hangman/game"
	"hangman/words"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <word_file>")
		os.Exit(1)
	}

	wordFile := os.Args[1]
	wordList, err := words.LoadWordsFromFile(wordFile)
	if err != nil {
		fmt.Printf("Error reading word file: %v\n", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	word := words.GetRandomWord(wordList)
	g := game.NewGame(word)

	fmt.Println("Welcome to Hangman!")
	fmt.Printf("The word to guess has %d letters.\n", len(word))

	scanner := bufio.NewScanner(os.Stdin)

	for !g.IsGameOver() {
		fmt.Println(game.GetHangmanArt(g.RemainingAttempts))
		fmt.Printf("Word: %s\n", g.GetDisplayWord())
		fmt.Printf("Incorrect letters: %s\n", g.GetIncorrectGuesses())
		fmt.Printf("Attempts left: %d\n", g.RemainingAttempts)
		fmt.Print("Enter a letter or word: ")

		scanner.Scan()
		input := strings.TrimSpace(strings.ToUpper(scanner.Text()))

		if len(input) == 0 {
			fmt.Println("Please enter a letter or word.")
			continue
		}

		if len(input) == 1 {
			if g.HasGuessed(rune(input[0])) {
				fmt.Println("You've already guessed that letter.")
				continue
			}
			correct := g.MakeGuess(rune(input[0]))
			if correct {
				fmt.Println("Good guess!")
			} else {
				fmt.Println("Incorrect guess.")
			}
		} else {
			if input == g.Word {
				g.RevealWord()
				break
			} else {
				fmt.Println("That's not the correct word.")
				g.RemainingAttempts -= 2
				if g.RemainingAttempts < 0 {
					g.RemainingAttempts = 0
				}
			}
		}
	}

	fmt.Println(game.GetHangmanArt(g.RemainingAttempts))
	if g.IsWon() {
		fmt.Println("Congratulations! You've won!")
	} else {
		fmt.Printf("Game over. The word was: %s\n", g.Word)
	}
}
