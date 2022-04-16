package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	listStories := make([]topstories, 0)
	result := TopStoriesGet()

	if CheckTime() {
		const basePath = "templates"

		router := http.NewServeMux()
		router.HandleFunc("/api/top", func(writer http.ResponseWriter, request *http.Request) {
			templates := template.Must(template.ParseFiles(basePath + "/_layout.html"))
			err := templates.Execute(writer, result)
			if err != nil {
				return
			}
		})

		err := http.ListenAndServe(":9000", router)
		checkError(err)
	}

	for _, ins := range result {
		listStories = append(listStories, *ins)
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
}
