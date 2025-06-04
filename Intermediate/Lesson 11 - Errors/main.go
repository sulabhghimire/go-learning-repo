package main

import (
	"errors"
	"fmt"
	"math"
)

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: square root of a negative number")
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty data")
	}
	return nil
}

func main() {

	result, err := squareRoot(16)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result1, err1 := squareRoot(-16)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	data := []byte{}
	if err = process(data); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data processed sucessfully")
	}

	if err = eProcess(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data processed sucessfully")
	}

	if err = readData(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config read sucessfully")
	}

}

type myError struct {
	message string
}

func (m myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eProcess() error {
	return &myError{"Custom error process"}
}

func readConfig() error {
	return errors.New("config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}
