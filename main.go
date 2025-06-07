package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
	scoreWPM    int
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
		case "Profile":
			// Add profile functionality later
			pterm.Info.Println("Profile not implemented yet")
		case "Exit":
			os.Exit(0)
		}
	}
}

// Display the Home
func displayHome() string {
	selectedDef, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"Easy", "Medium", "Hard"}).Show("Select Difficulty")

	var wordCount int
	switch selectedDef {
	case "Easy":
		wordCount = 10
	case "Medium":
		wordCount = 15
	case "Hard":
		wordCount = 20
	}

	value := selectRandomWords(wordCount)
	GameStart(value)
	return selectedDef
}

// Select Random words
func selectRandomWords(def int) []string {
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
	return words[:def]
}

func GameStart(words []string) {
	// Set game duration based on difficulty
	var gameDuration time.Duration
	switch len(words) {
	case 10:
		gameDuration = 60 * time.Second // Easy
	case 15:
		gameDuration = 90 * time.Second // Medium
	case 20:
		gameDuration = 120 * time.Second // Hard
	}

	// Start timer
	startTime := time.Now()
	endTime := startTime.Add(gameDuration)
	scanner := bufio.NewScanner(os.Stdin)

	// Clear screen
	pterm.Print("\033[H\033[2J") // ANSI escape code to clear screen

	for i, word := range words {
		// Check if time is up
		if time.Now().After(endTime) {
			pterm.Error.Println("Time's up!")
			fmt.Printf("You completed %d/%d words\n", i, len(words))
			return
		}

		// Display current word and timer
		remainingTime := time.Until(endTime).Round(time.Second)
		pterm.Info.Printf("Time left: %v\n", remainingTime)
		pterm.DefaultHeader.WithFullWidth(true).Println("Type the word:")
		fmt.Println(word)

		// Wait for user to press Enter
		fmt.Print("Your input: ")
		scanner.Scan() // Wait for Enter key

		// Clear screen for next word
		pterm.Print("\033[H\033[2J")
	}

	// Game completed
	elapsedTime := time.Since(startTime).Seconds()
	wpm := float64(len(words)) / (elapsedTime / 60.0)
	pterm.Success.Println("Game Completed!")
	fmt.Printf("Words typed: %d\n", len(words))
	fmt.Printf("Time taken: %.2f seconds\n", elapsedTime)
	fmt.Printf("Words per minute (WPM): %.2f\n", wpm)
}
