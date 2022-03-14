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
		fmt.Print("foo")
		cp.Lock()
		time.Sleep(10 * time.Millisecond)
		cp.Unlock()
	}
}

func (cp *ConcurrentPrinter) PrintBar(times int) {

	for i := 0; i < times; i++ {
		fmt.Print("bar")
		cp.Lock()
		time.Sleep(10 * time.Millisecond)
		cp.Unlock()
	}

}

func main() {

	times := 10
	cp := &ConcurrentPrinter{}

	cp.Add(2)

	go func() {
		cp.PrintBar(times)
		cp.Done()
	}()

	go func() {
		cp.PrintFoo(times)
		cp.Done()
	}()

	cp.Wait()
}
