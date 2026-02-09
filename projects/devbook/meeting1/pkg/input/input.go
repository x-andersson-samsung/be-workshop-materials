package input

import (
	"bufio"
	"fmt"
	"io"
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

	answer, _ := c.readLine()

	var optionNum int64
	optionNum, err = strconv.ParseInt(answer, 10, 64)
	if err != nil {
		fmt.Println("Invalid option number")
		return c.AskForOption(question, options)
	}

	return int(optionNum), nil
}
