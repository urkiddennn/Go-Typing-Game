package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

const historyFile = "records.json"

// For getting the current time and date
type (
	TimeStamp time.Time
	Highest   struct {
		highestWPM int
	}
)

type currentScore struct {
	id         int     `json:"id"`
	scoreWPM   float64 `json:"score_wpm"`
	Created_at string  `json:"created_at"`
}

type typeHistory struct {
	SCOREWPM []currentScore `json:"score_wpm"`
}

func main() {
	pterm.Info.Println("Select Menu")

	th := loadHistory()

	for {
		choice, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"Home", "Profile", "Exit"}).Show("Select Where you want to go")
		switch choice {
		case "Home":
			displayHome(&th)
		case "Profile":
			// Add profile functionality later
			profilePage(th)
		case "Exit":
			os.Exit(0)
		}
	}
}

// Display the Home
func displayHome(th *typeHistory) string {
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
	GameStart(value, th)
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

// Game start state
func GameStart(words []string, th *typeHistory) {
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
	saveRecord(th, wpm)

	fmt.Printf("Words typed: %d\n", len(words))
	fmt.Printf("Time taken: %.2f seconds\n", elapsedTime)
	fmt.Printf("Words per minute (WPM): %.2f\n", wpm)
}

// Profile Page
func profilePage(th typeHistory) {
	pterm.Info.Println("This is your profile")

	//	if len(th.SCOREWPM) == 0 {
	//	pterm.Info.Println("No available history")
	//	return
	//}
	table := pterm.TableData{{"ID", "Score WPM", "Date"}}
	for _, typeH := range th.SCOREWPM {
		wpmStr := strconv.FormatFloat(typeH.scoreWPM, 'f', 2, 64)
		ID := strconv.Itoa(typeH.id)
		table = append(table, []string{ID, wpmStr, typeH.Created_at})
	}
	pterm.DefaultTable.WithHasHeader().WithData(table).Render()
}

// Save records
func saveRecord(th *typeHistory, wpm float64) {
	id := len(th.SCOREWPM) + 1
	createdAt := time.Now().Format(time.RFC3339)
	crScore := currentScore{
		id:         id,
		scoreWPM:   wpm,
		Created_at: createdAt,
	}
	th.SCOREWPM = append(th.SCOREWPM, crScore)

	file, err := os.Create(historyFile)
	if err != nil {
		pterm.Error.Printf("Error saving Records: %v\n", err)
	}
	defer file.Close()
	pterm.Success.Println("Score added to history!")
}

func loadHistory() typeHistory {
	var th typeHistory

	file, err := os.Open(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return th
		}
		pterm.Error.Printf("Errpr loading Records: %v\n ", err)
		return th
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&th); err != nil {
		pterm.Error.Printf("Errr decoding history: %v\n", err)
		return th
	}

	return th
}
