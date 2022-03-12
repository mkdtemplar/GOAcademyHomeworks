package main

import "fmt"

type Action func() error

func SafeExec(a Action) Action {
	return func() error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Errorf("Recovered in function: %s", r)

			}
		}()
		panic("BOOM")
	}
}

func main() {

}
