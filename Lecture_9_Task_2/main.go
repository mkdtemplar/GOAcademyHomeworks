package main

import (
	"fmt"
	"math"
)

type Action func() error

func (a Action) Error() string {
	var aa float64
	scan, err := fmt.Scan(&aa)
	if err != nil {
		fmt.Println(math.Sqrt(float64(scan)))
		return "Error"
	} else {

		return "OK"
	}
	panic("implement me")
}

func (a Action) Divide() float64 {

	return 1 / 0
}

func SafeExec(a Action) Action {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("recovered in function: %s", r)

		}
	}()

	a.Divide()
	panic("BOOM")
}

func main() {

	var ar Action
	fmt.Println(SafeExec(ar).Divide())

}
