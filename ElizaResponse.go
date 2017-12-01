package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func elizaResponse(userInput string) string {
	input := userInput
	pattern := `(?i).*\bfather\b.*`
	pattern2 := `(?i)I am ([^.?!]*)[.?!]?`
	output := "why dont you tell me more about your father?"
	response := ""
	rand.Seed(time.Now().UTC().UnixNano())

	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	re := regexp.MustCompile(pattern2)
	//check if pattern is fouund in string (father/I am)
	if matched, _ := regexp.MatchString(pattern, input); matched {
		response = "User :" + input + "\nEliza :" + output
	} else if re.MatchString(input) {
		match := re.ReplaceAllString(input, "How do you know you are $1?")
		response = "User :" + input + " \nEliza : " + match
	} else {
		response = "User :" + input + " \nEliza :" + answers[rand.Intn(len(answers))]
	}
	return response
} //elizaResponse

func main() {
	// inputs

	userInput := []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I'm looking forward to the weekend.",
		"My grandfather was French!",
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?",
	}

	rand.Seed(time.Now().UTC().UnixNano())

	word := elizaResponse(userInput[rand.Intn(len("userInput"))])
	fmt.Print(word)
}
