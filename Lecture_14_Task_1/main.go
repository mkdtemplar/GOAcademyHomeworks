package main

import (
	"fmt"
	"net/http"
	"os"
)

func fetch(url string, ch chan<- string) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("Got response for %s: %d\n", url, resp.StatusCode)
}

func main() {

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}
