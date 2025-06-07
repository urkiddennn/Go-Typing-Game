package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pterm/pterm"
)

// For getting the current time and date
type (
	TimeStamp time.Time
	Highest   struct {
		highestWPM int
		HighestAcc float64
	}
)

type typeHistory struct {
	id          int
	scoreWMP    int
	accuracyWPM float64
	Created_at  *TimeStamp
}

func main() {
	pterm.Info.Println("Select Menu")

	for {
		choice, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"Home", "Profile", "Exit"}).Show("Select Where you want to go")
		switch choice {
		case "Home":
			displayHome()
		}
	}
}

// Display the Home
func displayHome() string {
	selectedDef, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"Easy", "Medium", "Hard"}).Show("Select Difficulty")

	switch selectedDef {
	case "Easy":
		value := selectRandomWords(10)
		GameStart(value)
	case "Medium":
		selectRandomWords(15)

	case "Hard":
		selectRandomWords(20)
	}

	return selectedDef
}

// Select Random words
func selectRandomWords(def int) []string {
	// easyLevelTime := 3 * time.Second

	// gameStart := false
	words := []string{
		"apple", "breeze", "cactus", "dolphin", "eagle",
		"fossil", "glacier", "horizon", "island", "jigsaw", "kitten", "lantern",
		"mosaic", "nectar", "oasis", "puzzle", "quartz", "river", "shadow",
		"tiger", "umbrella", "violet", "whisper", "xylophone", "yogurt",
		"zebra", "anchor", "blizzard", "canyon", "desert", "emerald",
		"flame", "guitar", "hammock", "indigo", "jungle", "kayak",
		"lotus", "meadow", "noodle", "orchid", "pebble", "quilt",
		"rainbow", "sapphire", "tulip", "vortex", "willow",
		"zenith", "compass", "are", "you", "and", "how", "welcome", "subscribe",
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	selectedRandomWords := words[:def]

	//	fmt.Println("Selected words", selectedRandomWords)
	return selectedRandomWords
}

func GameStart(words []string) {
	for i := range len(words) {
		time.Sleep(2 * time.Second)
		fmt.Println(words[i])
	}
}
