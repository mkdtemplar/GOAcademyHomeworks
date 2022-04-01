package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	buffer chan string
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	buff := make(chan string, bufferSize)

	newBufferedCTX := &BufferedContext{Context: ctx, buffer: buff, CancelFunc: cancel}

	return newBufferedCTX
}

func (bc *BufferedContext) Done() <-chan struct{} {
	/* This function will serve in place of the oriignal context */
	//make it so that the result channel gets closed in one of the to cases;
	// a) the emebdded context times out
	//b) the buffer gets filled

	if len(bc.buffer) == cap(bc.buffer) {
		fmt.Println("Buffer limit reached")
		bc.CancelFunc()
	}

	return bc.Context.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	/* This function serves for executing the test */
	/* Implement the rest */
	fn(bc, bc.buffer)
}

func main() {

	ctx := NewBufferedContext(time.Second, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("We are done TIME out")
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200)
				fmt.Println("bar")
			}
		}
	})
}
