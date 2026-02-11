package input

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type CLI struct {
	in     io.Reader
	out    io.Writer
	prompt string
}

func NewCLI(in io.Reader, out io.Writer) *CLI {
	return &CLI{
		in:     in,
		out:    out,
		prompt: "> ",
	}
}

func (c *CLI) readLine() (string, error) {
	scanner := bufio.NewScanner(c.in)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), scanner.Err()
}

func (c *CLI) AskForString(question string) (string, error) {
	_, err := fmt.Fprintf(c.out, "%s\n%s", question, c.prompt)
	if err != nil {
		return "", err
	}

	answer, _ := c.readLine()
	return strings.TrimSpace(answer), nil
}

func (c *CLI) AskForOption(question string, options []string) (int, error) {
	_, err := fmt.Fprintln(c.out, question)
	if err != nil {
		return -1, err
	}

	for idx, option := range options {
		if _, err = fmt.Fprintf(c.out, "%d: %s\n", idx, option); err != nil {
			return -1, err
		}
	}

	if _, err = fmt.Fprint(c.out, c.prompt); err != nil {
		return -1, err
	}

	answer, err := c.readLine()
	if err != nil {
		return -1, err
	}

	var optionNum int64
	optionNum, err = strconv.ParseInt(answer, 10, 64)
	if err != nil {
		fmt.Println("Invalid option number")
		return c.AskForOption(question, options)
	}

	return int(optionNum), nil
}

func (c *CLI) AskForMapOption(question string, options map[string]string) (string, error) {
	// Print question
	_, err := fmt.Fprintln(c.out, question)
	if err != nil {
		return "", err
	}

	// Check for longest key for formatting
	maxLength := 0
	for key := range options {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}

	// Sort by key
	type option struct {
		key   string
		value string
	}
	sortedOptions := make([]option, 0, len(options))
	for k, v := range options {
		sortedOptions = append(sortedOptions, option{key: k, value: v})
	}
	sort.Slice(sortedOptions, func(i, j int) bool { return sortedOptions[i].key < sortedOptions[j].key })

	// Print options
	for _, choice := range sortedOptions {
		if _, err = fmt.Fprintf(c.out, "%*s: %s\n", maxLength, choice.key, choice.value); err != nil {
			return "", err
		}
	}

	// Print prompt
	if _, err = fmt.Fprint(c.out, c.prompt); err != nil {
		return "", err
	}

	// Read answer
	answer, err := c.readLine()
	if err != nil {
		return "", err
	}

	// Validate
	if _, ok := options[answer]; !ok {
		fmt.Println("Invalid option")
		return c.AskForMapOption(question, options)
	}

	return answer, nil
}
