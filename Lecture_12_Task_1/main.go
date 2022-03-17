package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	
	channel := make(chan string)
	
	clearInterval = time.Duration(rand.Intn(1e3))
	
	var wg sync.WaitGroup
	
	wg.Add(bufferLimit)
	
	for i := 0; i < bufferLimit; i++ {
		go func() {
			channel <- data
			wg.Done()
			time.Sleep(clearInterval * time.Millisecond)
			close(channel)
		}()
	}
	
	return channel
}
func main() {

	clearInterval := time.Duration(rand.Intn(1e3))
	
	c := generateThrottled("Data string", 2, clearInterval)

	for i := range c {
		
		log.Println(i)
	}
}
