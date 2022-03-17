package main

import (
	"fmt"
)

func processEven(inputs []int) chan []int {

	even := make(chan []int)
	evenSlice := make([]int, 0)
	go func(e chan []int) {
		defer close(e)
		for i := range inputs {
			if inputs[i]%2 == 0 {
				evenSlice = append(evenSlice, inputs[i])
			}
		}
		even <- evenSlice
	}(even)

	return even

}
func processOdd(inputs []int) chan []int {

	odd := make(chan []int)
	oddSlice := make([]int, 0)
	go func(odd chan []int) {
		o := odd
		defer close(o)
		for i := range inputs {
			if inputs[i]%2 != 0 {
				oddSlice = append(oddSlice, inputs[i])

			}
		}
		o <- oddSlice
	}(odd)
	return odd
}

func main() {

	inputs := []int{1, 17, 34, 56, 2, 8}

	evenCh := processEven(inputs)
	oddCh := processOdd(inputs)

	for i := range oddCh {
		fmt.Println(i)

	}

	for j := range evenCh {
		fmt.Println(j)
	}
}
