package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		err1    = errors.New("someError")
		err2    = errors.New("someError")
		errWrap = fmt.Errorf("%s:%w", "module", err1)
	)

	fmt.Printf("%s | %s | %s\n", err1, err2, errWrap)
	fmt.Printf("%t | %t | %t\n", err1 == err2, errors.Is(err1, err2), err1.Error() == err2.Error())
	fmt.Printf("%t | %t | %t\n", err1 == errWrap, errors.Is(errWrap, err1), err1.Error() == errWrap.Error())
	fmt.Printf("%t | %t\n", errors.Is(errWrap, err1), errors.Is(err1, errWrap))
}
