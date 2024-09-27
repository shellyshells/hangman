package words

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func word_file() string {
	return "words.txt"
}

func LoadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func GetRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

var words = []string{
	"abruptly",
	"absurd",
	"abyss",
	"affix",
	"askew",
	"avenue",
	"awkward",
	"axiom",
	"azure",
	"bagpipes",
	"bandwagon",
	"banjo",
	"bayou",
	"beekeeper",
	"bikini",
	"blitz",
	"blizzard",
	"boggle",
	"bookworm",
	"boxcar",
	"boxful",
	"buckaroo",
	"buffalo",
	"buffoon",
	"buxom",
	"buzzard",
	"buzzing",
	"buzzwords",
	"caliph",
	"cobweb",
	"cockiness"}
