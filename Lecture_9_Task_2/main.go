package main

import (
	"errors"
	"fmt"
	"math"
)

type Action func() error

func SafeExec(a Action) Action {
	return func() error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Errorf("recovered in function: %s", r)

			}
		}()
		panic("BOOM")
	}
	return a
}

func Function() error {

	return errors.New("error")

}

func main() {

	var ar Action

	ar = func() float64 {
		return math.Sqrt(-1)
	}
	SafeExec(ar)
}
