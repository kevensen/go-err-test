package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

var (
	divisor   = flag.Int("divisor", 0, "The number doing the dividing")
	dividend  = flag.Int("dividend", 0, "The number being divided")
	ErrCustom = errors.New("custom error")
)

type DivideByZeroError struct {
	err error
}

func (d *DivideByZeroError) Error() string {
	return fmt.Sprintf("cannot divide by zero: %v", d.err)
}

func (d *DivideByZeroError) Unwrap() error {
	return ErrCustom
}

func (d *DivideByZeroError) Is(target error) bool {
	return target == &DivideByZeroError{}
}

func divide(dividend, divisor int) (quotient int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("division error: %w", &DivideByZeroError{r.(error)})
		}
	}()

	quotient = dividend / divisor

	return
}

func main() {
	q, err := divide(*dividend, *divisor)
	if err != nil {
		var dbz *DivideByZeroError
		if errors.As(err, &dbz) {
			fmt.Printf("one way to check a divide by zero error\n")
		}

		if errors.Is(err, ErrCustom) {
			fmt.Printf("yet another way to check a divide by zero error\n")
		}

		log.Fatal(err)
	}

	fmt.Println("Quotient: ", q)
}
