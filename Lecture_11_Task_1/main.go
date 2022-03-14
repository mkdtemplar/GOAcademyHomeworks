package main

import (
	"fmt"
	"sync"
)

type Concurency struct {
	sync.WaitGroup
	sync.Mutex
}

func processEven(inputs []int) chan int {

	even := make(chan int)
	for i := range inputs {
		if inputs[i]%2 == 0 {
			return even
		}
	}

	return nil

}
func processOdd(inputs []int) chan int {

	odd := make(chan int)
	for i := range inputs {
		if inputs[i]%2 != 0 {
			return odd
		}
	}
	return nil
}

func main() {

	inputs := []int{1, 17, 34, 56, 2, 8}
	cp := &Concurency{}

	cp.Add(2)

	go func() {
		fmt.Println(processEven(inputs))
		cp.Done()
	}()

	go func() {
		fmt.Println(processOdd(inputs))
		cp.Done()
	}()

	cp.Wait()

}
