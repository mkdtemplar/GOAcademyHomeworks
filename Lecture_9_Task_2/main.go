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
			if a := recover(); a != nil {
				fmt.Errorf("recovered in function: %s", a)

			}
		}()

		panic("BOOM")
	}

}

func Function() error {

	return errors.New("error")

}

func main() {

	var ar Action

	ar = func() error {
		var aa int
		scan, err := fmt.Scan(&aa)
		if err != nil {
			fmt.Println(math.Sqrt(float64(scan)))
			return err
		} else {

			return nil
		}
	}
	SafeExec(ar)
}
