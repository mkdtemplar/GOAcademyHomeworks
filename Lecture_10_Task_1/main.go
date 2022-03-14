package main

import (
	"fmt"
	"sync"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {

	for i := 0; i < times; i++ {

		fmt.Print("foo")
	}

}

func (cp *ConcurrentPrinter) PrintBar(times int) {

	for i := 0; i < times; i++ {

		fmt.Print("bar")

	}

}

func main() {

	times := 10
	cp := &ConcurrentPrinter{}

	cp.Add(2)

	go func() {
		cp.PrintFoo(times)
		//cp.Lock()
		cp.Done()
		//cp.Unlock()
	}()

	go func() {
		//cp.Lock()
		cp.PrintBar(times)

		cp.Done()
		//cp.Unlock()
	}()

	cp.Wait()
}
