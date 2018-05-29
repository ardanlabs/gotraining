// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a chat bot to talk to you. Read lines of input from stdin and respond
// with random responses to predefined keywords.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	// Print something to welcome users and tell them what they can ask about.
	fmt.Println(`Welcome to my chatbot! Exit with ctrl-c or type "exit" or "quit".`)
	fmt.Println("You can ask me about:")
	fmt.Println("\tThe weather.")
	fmt.Println("\tHow I'm feeling.")
	fmt.Println("\tThe game last night.")

	// Print a > to prompt the user to type something. Don't use Println for this.
	fmt.Print("> ")

	// Create a bufio.Scanner that will read lines from os.Stdin.
	s := bufio.NewScanner(os.Stdin)

	// Call the Scan method of the scanner until it returns false.
	for s.Scan() {

		// Get the user's input from the Text() method.
		input := strings.ToLower(s.Text())

		// Decide if you should exit.
		if input == "exit" || input == "quit" {
			fmt.Println("< Goodbye! Thanks for chatting!")
			break
		}

		// Choose a response based on their input.
		response := answer(input)

		// Print a < followed by the chatbot's response. Try adding a short random
		// delay between characters so it looks like someone is typing.
		fmt.Print("< ")
		for _, r := range response {
			delay := rand.Intn(100) + 1
			time.Sleep(time.Duration(delay) * time.Millisecond)
			fmt.Print(string(r))
		}

		// Start the next line and print a > again to prompt for more input.
		fmt.Print("\n> ")
	}

	// Check the Err() method of the scanner to see if it failed for some reason.
	if err := s.Err(); err != nil {
		log.Fatal("could not read", err)
	}
}

var weather = []string{
	"Not bad, not bad!",
	"Raining now but it should let up soon!",
	"It's looking pretty dreary :(",
}

var mood = []string{
	"Never better!",
	"Kind of sleepy.",
	"I could use a hug.",
}

var game = []string{
	"What was Wenger thinking sending Walcott on that early?",
	"See, the thing about Arsenal is they always try to walk it in!",
}

func answer(question string) string {
	switch {
	case strings.Contains(question, "weather"):
		return weather[rand.Intn(len(weather))]
	case strings.Contains(question, "feeling"):
		return mood[rand.Intn(len(mood))]
	case strings.Contains(question, "game") || strings.Contains(question, "ludicrous display"):
		return game[rand.Intn(len(game))]
	}
	return "Sorry, I didn't get that. Try asking me about the *weather*, how I'm *feeling*, or the *game* last night."
}
