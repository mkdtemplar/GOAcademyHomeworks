package main

import (
	"Lecture_27_Task/sqlcCode"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

type topstories struct {
	STORY_ID  int
	TITLE     string
	SCORE     int
	URL       string
	TimeStamp string
}

type StoryId []int

var wg = &sync.WaitGroup{}

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
	conn, err := sql.Open("sqlite", "Stories.db")
	checkError(err)
	db := sqlcCode.New(conn)

	timeDb, err := db.GetTimeFromDB(context.Background())
	checkError(err)

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
	return false
}
func TopStoriesGet() []*topstories {

	conn, err := sql.Open("sqlite", "Stories.db")
	checkError(err)
	db := sqlcCode.New(conn)
	data := GetStoryID()

	ts := make([]*topstories, 0)
	db.DeleteAllRecords(context.Background())
	for ids := 0; ids < 20; ids++ {
		wg.Add(1)
		go func(res []*topstories, wg *sync.WaitGroup) {
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

			var responseStories = topstories{}
			err = json.Unmarshal(body, &responseStories)
			if err != nil {
				return
			}

			responseStories = topstories{STORY_ID: data[ids], SCORE: responseStories.SCORE, TITLE: responseStories.TITLE,
				URL: responseStories.URL, TimeStamp: time.Now().Format("2006-01-02 15:04:05")}

			db.InsertData(context.Background(), sqlcCode.InsertDataParams{
				StoryID:   int64(data[ids]),
				Title:     responseStories.TITLE,
				Score:     int64(responseStories.SCORE),
				Url:       responseStories.URL,
				Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			})

			ts = append(ts, &responseStories)

			wg.Done()
		}(ts, wg)

		wg.Wait()
	}

	return ts
}
