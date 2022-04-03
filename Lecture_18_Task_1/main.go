package main

import (
	"fmt"
	"sync"
	"time"
)

func goPrimesAndSleep(n int, sleep time.Duration) []int {
	var res []int
	var wg = &sync.WaitGroup{}

	for k := 2; k < n; k++ {
		wg.Add(1)
		go func(result []int, wg *sync.WaitGroup) {
			for i := 2; i < n; i++ {
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i {
						res = append(res, k)
					}
					break
				}
			}
			wg.Done()
		}(res, wg)
		wg.Wait()
	}

	return res
}

func main() {
	fmt.Println(goPrimesAndSleep(100, time.Millisecond))
}
