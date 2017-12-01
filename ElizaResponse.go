package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func elizaResponse(userInput string) string {
	input := userInput
	pattern := `(?i).*\bFather\b.*`
	output := "Why don’t you tell me more about your father?"
	response := ""
	rand.Seed(time.Now().UTC().UnixNano())
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	//check if pattern is fouund in string
	if matched, _ := regexp.MatchString(pattern, input); matched {
		response = "User :" + input + " \nEliza :" + output
	} else {
		response = "User :" + input + " \nEliza :" + answers[rand.Intn(len(answers))]
	}
	return response
} //elizaResponse

func main() {
	//slice of inputs
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
