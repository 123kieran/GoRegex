package main

import (
	"fmt"
	"math/rand"
	"time"
)

func elizaResponse(inputStr string) string {
	//return random string output
	input := inputStr
	rand.Seed(time.Now().UTC().UnixNano())
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	response := "User :" + input + " \n Eliza :" + answers[rand.Intn(len(answers))]
	return response
} //elizaResponse

func main() {
	// inputs put into array
	userInput := []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I'm looking forward to the weekend.",
		"My grandfather was French!",
	}
	rand.Seed(time.Now().UTC().UnixNano())

	word := elizaResponse(userInput[rand.Intn(len(userInput))])

	fmt.Print(word)
}
