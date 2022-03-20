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
	for i := 0; i < bufferSize; i++ {
		go func() {
			done := bc.Done()
			defer cancel()
			for {
				select {
				case <-done:
					return
				case ch := <-chanel:
					fmt.Println(ch)
				}
				return
			}
			close(chanel)
		}()
		wg.Done()
	}

	return &bc
}
func (bc *BufferedContext) Done() <-chan struct{} {
	/* This function will serve in place of the oriignal context */
	/* make it so that the result channel gets closed in one of the to cases;
	   a) the emebdded context times out
	   b) the buffer gets filled
	*/

	return bc.Context.Done()
}
func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	/* This function serves for executing the test */
	/* Implement the rest */
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	chanel := make(chan string)

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

/*
unc NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {

    ctx, cancel := context.WithTimeout(context.Background(), timeout)



    defer cancel()

    bc := &BufferedContext{ctx, 10, timeout}



    return bc



}

func (bc *BufferedContext) Done() <-chan struct{} {



    ctx, cancel := context.WithTimeout(bc, bc.timeout)

    defer cancel()

    return ctx.Done()

}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {

    ch := make(chan string, bc.buffer)

    for len(ch) == cap(ch) {

        ch <- "bar"

        time.Sleep(time.Millisecond * 200)

    }

    fn(bc.Context, ch)
}
*/
