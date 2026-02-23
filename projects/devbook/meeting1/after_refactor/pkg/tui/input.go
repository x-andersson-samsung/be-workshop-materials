package tui

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const prompt = "> "

func readLine() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), scanner.Err()
}

func AskForString(question string) (string, error) {
	_, err := fmt.Printf("%s\n%s", question, prompt)
	if err != nil {
		return "", err
	}

	answer, _ := readLine()
	return strings.TrimSpace(answer), nil
}

func AskForMapChoice(question string, choices map[string]string) (string, error) {
	// Print question
	_, err := fmt.Println(question)
	if err != nil {
		return "", err
	}

	// Check for longest key for formatting
	maxLength := 0
	for key := range choices {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}

	// Sort by key
	type choiceStruct struct {
		key   string
		value string
	}
	sortedChoices := make([]choiceStruct, 0, len(choices))
	for k, v := range choices {
		sortedChoices = append(sortedChoices, choiceStruct{key: k, value: v})
	}
	sort.Slice(sortedChoices, func(i, j int) bool { return sortedChoices[i].key < sortedChoices[j].key })

	// Print choices
	for _, choice := range sortedChoices {
		if _, err = fmt.Printf("%*s: %s\n", maxLength, choice.key, choice.value); err != nil {
			return "", err
		}
	}

	// Print prompt
	if _, err = fmt.Print(prompt); err != nil {
		return "", err
	}

	// Read answer
	answer, err := readLine()
	if err != nil {
		return "", err
	}

	// Validate
	if _, ok := choices[answer]; !ok {
		fmt.Println("Invalid choice")
		return AskForMapChoice(question, choices)
	}

	return answer, nil
}
