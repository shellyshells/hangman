package game

var hangmanArt = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

func GetHangmanArt(remainingAttempts int) string {
	index := 6 - remainingAttempts
	if index < 0 {
		index = 0
	} else if index >= len(hangmanArt) {
		index = len(hangmanArt) - 1
	}
	return hangmanArt[index]
}
