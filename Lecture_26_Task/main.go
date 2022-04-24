package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	var stories []topstories
	var err error
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

		go func() {
			log.Fatal(http.ListenAndServe(":9000", router))
			exit <- struct{}{}
		}()
		switch runtime.GOOS {
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:9000/api/top").Start()
		case "darwin":
			err = exec.Command("open", "http://localhost:9000/api/top").Start()
		case "linux":
			err = exec.Command("xdg-open", "http://localhost:9000/api/top").Start()
		}
		checkError(err)
		<-exit
	} else if CheckTime() == false {
		db, err := gorm.Open(sqlite.Open("Stories.db"), &gorm.Config{})
		checkError(err)
		db.Find(&stories)

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
			jm, err := json.MarshalIndent(stories, "", " ")
			checkError(err)
			fmt.Println(string(jm))
			break
		case 2:
			for _, st := range stories {
				fmt.Println("Story Id: ", st.STORY_ID)
				fmt.Println("Title: ", st.TITLE)
				fmt.Println("Score: ", st.SCORE)
				fmt.Println("URL: ", st.URL)
				fmt.Println("Time stamp: ", st.TimeStamp)
			}
			break
		}
	}

}
