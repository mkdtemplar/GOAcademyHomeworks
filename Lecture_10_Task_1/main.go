package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {

	for i := 0; i < times; i++ {
		cp.Lock()
		fmt.Print("foo")
		defer cp.Unlock()
	}

}

func (cp *ConcurrentPrinter) PrintBar(times int) {

	time.Sleep(10 * time.Millisecond)
	for i := 0; i < times; i++ {
		cp.Lock()
		fmt.Print("bar")
		defer cp.Unlock()
	}

}

func main() {

	times := 10
	cp := &ConcurrentPrinter{}

	cp.Add(2)

	go func() {
		cp.PrintFoo(times)
		cp.Done()
	}()

	go func() {
		cp.PrintBar(times)
		cp.Done()
	}()

	cp.Wait()
}
