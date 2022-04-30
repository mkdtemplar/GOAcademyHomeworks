package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type TopStories struct {
	Score int    `json:"score"`
	Title string `json:"title"`
	Url   string `json:"url"`
}
type StoryID []int

var wg = &sync.WaitGroup{}

func GetStoryID() StoryID {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	checkError(err)

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)

	var response StoryID

	err = json.Unmarshal(body, &response)

	return response
}

func TopStoriesGet() *[]TopStories {

	data := GetStoryID()
	ts := make([]TopStories, 0)

	for ids := 0; ids < 10; ids++ {
		wg.Add(1)
		go func(res *[]TopStories, wg *sync.WaitGroup) {
			url := fmt.Sprintf("%s%d%s", "https://hacker-news.firebaseio.com/v0/item/", data[ids], ".json?print=pretty")

			req, err := http.Get(url)
			checkError(err)

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(req.Body)

			body, err := ioutil.ReadAll(req.Body)
			checkError(err)

			var response = TopStories{}
			err = json.Unmarshal(body, &response)
			if err != nil {
				return
			}

			response = TopStories{Score: response.Score, Title: response.Title, Url: response.Url}

			checkError(err)
			ts = append(ts, response)
			wg.Done()
		}(&ts, wg)

		wg.Wait()
	}
	return &ts
}

func main() {

	result := TopStoriesGet()

	const basePath = "templates"

	router := http.NewServeMux()
	router.HandleFunc("/api/top", func(writer http.ResponseWriter, request *http.Request) {
		templates := template.Must(template.ParseFiles(basePath + "/_layout.html"))
		err := templates.Execute(writer, *result)
		if err != nil {
			return
		}
	})

	exit := make(chan struct{}, 1)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", router))
		exit <- struct{}{}
	}()

	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:8080/api/top").Start()
	if err != nil {
		fmt.Println(err)
	}
	<-exit
}
