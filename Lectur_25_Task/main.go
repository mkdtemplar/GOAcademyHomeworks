package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var wg = &sync.WaitGroup{}

type StoryId []int

type TopStories struct {
	Score int    `json:"score"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func GetStoryID() StoryId {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	checkError(err)

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)

	var response StoryId

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}
	return response
}

func TopStoriesGet() *[]TopStories {

	data := GetStoryID()

	ts := make([]TopStories, 0)

	for ids := 0; ids < 20; ids++ {
		wg.Add(1)
		go func(res []TopStories, wg *sync.WaitGroup) {
			url := fmt.Sprintf("%s%d%s", "https://hacker-news.firebaseio.com/v0/item/", data[ids], ".json?print=pretty")
			req, err1 := http.Get(url)
			checkError(err1)

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(req.Body)

			body, err := ioutil.ReadAll(req.Body)
			checkError(err)

			var responseStories = TopStories{}
			err = json.Unmarshal(body, &responseStories)
			if err != nil {
				return
			}

			responseStories = TopStories{Score: responseStories.Score, Title: responseStories.Title, Url: responseStories.Url}

			checkError(err)
			ts = append(ts, responseStories)

			wg.Done()
		}(ts, wg)

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
		err := templates.Execute(writer, result)
		if err != nil {
			return
		}
	})
	log.Fatal(http.ListenAndServe(":9000", router))

}
