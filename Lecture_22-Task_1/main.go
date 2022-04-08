package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type StoriesIDs struct {
	StoryID []int `json:"storyID"`
}

type TopStories struct {
	Score int    `json:"score"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

var wg = &sync.WaitGroup{}

func TopStoriesGet(file []byte, data StoriesIDs, ts []TopStories) *[]TopStories {

	//var resTs []byte
	data = StoriesIDs{}
	ts = make([]TopStories, len(data.StoryID))

	json.Unmarshal(file, &data)

	for ids := 0; ids < 10; ids++ {
		wg.Add(1)
		go func(res []TopStories, wg *sync.WaitGroup) {
			url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(data.StoryID[ids]) + ".json?print=pretty"

			req, err := http.Get(url)
			checkError(err)

			defer req.Body.Close()

			body, err := ioutil.ReadAll(req.Body)
			checkError(err)

			var response = TopStories{}
			json.Unmarshal(body, &response)

			response = TopStories{
				Score: response.Score,
				Title: response.Title,
				Url:   response.Url,
			}

			checkError(err)
			ts = append(ts, response)
			wg.Done()
		}(ts, wg)

		wg.Wait()
	}
	return &ts
}

func main() {

	file, err := ioutil.ReadFile("StoriesID.json")
	checkError(err)

	data := StoriesIDs{}
	ts := make([]TopStories, len(data.StoryID))

	result := TopStoriesGet(file, data, ts)

	const basePath = "templates"

	router := http.NewServeMux()
	router.HandleFunc("/api/top", func(writer http.ResponseWriter, request *http.Request) {
		templates := template.Must(template.ParseFiles(basePath + "/_layout.html"))
		templates.Execute(writer, result)
	})
	log.Fatal(http.ListenAndServe(":9000", router))
}
