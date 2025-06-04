package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	err     error
}

func (c *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", c.code, c.message, c.err)
}

// Function that returns a custom error
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Something went wrong!",
// 	}
// }

func doSomething() error {
	err := doSomethingAgain()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			err:     err,
		}
	}
	return nil
}

func doSomethingAgain() error {
	return errors.New("internal error")
}

func main() {

	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	} else {
		fmt.Println("Sucessfull")
	}

	// wrapped erros

}
