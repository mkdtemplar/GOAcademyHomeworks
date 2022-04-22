package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
	"sync"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var wg = &sync.WaitGroup{}
var db *sql.DB

type StoryId []int

type TopStories struct {
	Score int    `json:"score"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type DbStory struct {
	STORY_ID  int
	TITLE     string
	SCORE     int
	URL       string
	TimeStamp string
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

func CheckTime() bool {
	var list DbStory
	var listTimes []string
	db, err := sql.Open("sqlite", "Stories.db")
	row, err := db.Query("SELECT DISTINCT TimeStamp FROM topstories ORDER BY TimeStamp ")
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {

		}
	}(row)

	for row.Next() {
		err = row.Scan(&list.TimeStamp)
		listTimes = append(listTimes, list.TimeStamp)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(listTimes) > 0 {
		timeDb := listTimes[0]

		timeParsedFromDB, err := time.Parse("2006-01-02 15:04:05", timeDb)
		checkError(err)
		dateDb := time.Date(timeParsedFromDB.Year(), timeParsedFromDB.Month(), timeParsedFromDB.Day(),
			timeParsedFromDB.Hour(), timeParsedFromDB.Minute(), timeParsedFromDB.Second(), timeParsedFromDB.Nanosecond(), time.UTC)

		timeAccess := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(),
			time.Now().Second(), 0, time.UTC)

		if dateDb.Before(timeAccess) {
			timeHour := time.Now().Hour()
			timeParsedHour := timeParsedFromDB.Hour()
			if timeHour-timeParsedHour > 1 {
				return true
			}
		}
	}
	db.Close()
	return false
}

func TopStoriesGet() *[]TopStories {

	db, err := sql.Open("sqlite", "Stories.db")
	timeNow := time.Now().Format("2006-01-02 15:04:05")

	InsertSQL := `INSERT INTO topstories (STORY_ID, TITLE, SCORE, URL, TimeStamp) VALUES (?,?,?,?, ?)`

	statement, err := db.Prepare(InsertSQL)

	checkError(err)

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
			_, err2 := statement.Exec(data[ids], responseStories.Title, responseStories.Score, responseStories.Url, timeNow)
			checkError(err2)
			responseStories = TopStories{Score: responseStories.Score, Title: responseStories.Title, Url: responseStories.Url}

			checkError(err)
			ts = append(ts, responseStories)

			wg.Done()
		}(ts, wg)

		wg.Wait()
	}
	db.Close()
	return &ts
}

func main() {

	var list DbStory

	listStories := make([]DbStory, 0)

	db, err := sql.Open("sqlite", "Stories.db")

	if CheckTime() {
		DeleteSQL := `DELETE FROM topstories`

		statementDel, err := db.Prepare(DeleteSQL)
		checkError(err)

		_, err = statementDel.Exec()
		checkError(err)
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

	rowSelect, err := db.Query("SELECT * FROM topstories")
	checkError(err)
	defer func(rowSelect *sql.Rows) {
		err := rowSelect.Close()
		if err != nil {

		}
	}(rowSelect)
	for rowSelect.Next() {
		err = rowSelect.Scan(&list.STORY_ID, &list.TITLE, &list.SCORE, &list.URL, &list.TimeStamp)
		checkError(err)

		listStories = append(listStories, DbStory{
			STORY_ID: list.STORY_ID, TITLE: list.TITLE, SCORE: list.SCORE, URL: list.URL, TimeStamp: list.TimeStamp,
		})

	}

	var choice int

	fmt.Print("Enter 1. if you want to see results in JSON format or Enter 2 for regular format:")
	fmt.Scan(&choice)

	for choice != 1 && choice != 2 {
		fmt.Println("Wrong choice")
		fmt.Print("Enter 1. if you want to see results in JSON format or Enter 2 for regular format:")
		fmt.Scan(&choice)
	}
	switch choice {
	case 1:
		jm, err := json.MarshalIndent(listStories, "", " ")
		checkError(err)
		fmt.Println(string(jm))
		break
	case 2:
		for _, st := range listStories {
			fmt.Println("Story Id: ", st.STORY_ID)
			fmt.Println("Title: ", st.TITLE)
			fmt.Println("Score: ", st.SCORE)
			fmt.Println("URL: ", st.URL)
			fmt.Println("Time stamp: ", st.TimeStamp)
		}
		break
	}

	db.Close()
}
