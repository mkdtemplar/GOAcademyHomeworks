package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type BufferedContext struct {
	context.Context
	/* Add other fields you might need */
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	chanel := make(chan string, bufferSize)
	var wg sync.WaitGroup

	bc := BufferedContext{ctx}

	wg.Add(bufferSize)
	go func() {
		defer wg.Done()
		defer cancel()
		for {
			select {
			case <-bc.Done():
				return
			case ch := <-chanel:
				fmt.Println(ch)
			}
		}
	}()

	return &bc
}
func (bc *BufferedContext) Done() <-chan struct{} {
	/* This function will serve in place of the oriignal context */
	/* make it so that the result channel gets closed in one of the to cases;
	   a) the emebdded context times out
	   b) the buffer gets filled
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	bc = &BufferedContext{ctx}

	return ctx.Done()
}
func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	/* This function serves for executing the test */
	/* Implement the rest */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	chanel := make(chan string)
	bc = &BufferedContext{ctx}

	for i := 0; i < len(chanel); i++ {
		chanel <- "bar"
	}

	fn(ctx, chanel)
}

func main() {

	ctx := NewBufferedContext(time.Second, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200)
				fmt.Println("bar")
			}
		}
	})
}
