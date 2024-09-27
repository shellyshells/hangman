package game

import (
	"math/rand"
	"strings"
)

type Game struct {
	Word              string
	GuessedLetters    map[rune]bool
	RemainingAttempts int
	RevealedLetters   int
}

func NewGame(word string) *Game {
	g := &Game{
		Word:              strings.ToUpper(word),
		GuessedLetters:    make(map[rune]bool),
		RemainingAttempts: 6,
	}
	g.revealRandomLetters()
	return g
}

func (g *Game) revealRandomLetters() {
	toReveal := len(g.Word)/2 - 1
	for i := 0; i < toReveal; i++ {
		idx := rand.Intn(len(g.Word))
		g.GuessedLetters[rune(g.Word[idx])] = true
		g.RevealedLetters++
	}
}

func (g *Game) MakeGuess(letter rune) bool {
	g.GuessedLetters[letter] = true

	if strings.ContainsRune(g.Word, letter) {
		return true
	}

	g.RemainingAttempts--
	return false
}

func (g *Game) GetDisplayWord() string {
	var display strings.Builder
	for _, char := range g.Word {
		if g.GuessedLetters[char] {
			display.WriteRune(char)
		} else {
			display.WriteRune('_')
		}
		display.WriteRune(' ')
	}
	return display.String()
}

func (g *Game) GetIncorrectGuesses() string {
	var incorrect strings.Builder
	for letter := range g.GuessedLetters {
		if !strings.ContainsRune(g.Word, letter) {
			incorrect.WriteRune(letter)
			incorrect.WriteRune(' ')
		}
	}
	return incorrect.String()
}

func (g *Game) IsGameOver() bool {
	return g.IsWon() || g.RemainingAttempts <= 0
}

func (g *Game) IsWon() bool {
	for _, char := range g.Word {
		if !g.GuessedLetters[char] {
			return false
		}
	}
	return true
}

func (g *Game) HasGuessed(letter rune) bool {
	return g.GuessedLetters[letter]
}

func (g *Game) RevealWord() {
	for _, char := range g.Word {
		g.GuessedLetters[char] = true
	}
}
