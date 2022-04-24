package main

import (
	"Lecture_28_Task/sqlcCode"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	_ "modernc.org/sqlite"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	listStories := make([]sqlcCode.Topstory, 0)

	if CheckTime() == true {
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

		exit := make(chan struct{}, 1)
		// start the server
		go func() {
			fmt.Println("Listening on localhost:8080")
			http.ListenAndServe(":8080", router)
			exit <- struct{}{}
		}()

		var err error
		switch runtime.GOOS {
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:8080/api/top").Start()
		case "darwin":
			err = exec.Command("open", "http://localhost:8080/api/top").Start()
		case "linux":
			err = exec.Command("xdg-open", "http://localhost:8080/api/top").Start()
		}
		if err != nil {
			fmt.Println(err)
		}
		<-exit

	} else if CheckTime() == false {
		conn, err := sql.Open("sqlite", "Stories.db")
		checkError(err)
		db := sqlcCode.New(conn)
		result, err := db.ListAllStories(context.Background())
		checkError(err)
		for _, ins := range result {
			listStories = append(listStories, ins)
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
				fmt.Println("Story Id: ", st.StoryID)
				fmt.Println("Title: ", st.Title)
				fmt.Println("Score: ", st.Score)
				fmt.Println("URL: ", st.Url)
				fmt.Println("Time stamp: ", st.Timestamp)
			}
			break
		}
	}

}
