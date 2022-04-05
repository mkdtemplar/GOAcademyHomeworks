package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Drink struct {
	Name   string `json:"strDrink"`
	Recipe string `json:"StrInstructions"`
}

type Response struct {
	Drinks []Drink
}

func Start(drinkName string) string {
	return drinkName
}

func main() {

	fmt.Print("Enter drink name or enter 'nothing' to exit: ")
	drinkName := ""
	fmt.Scanln(&drinkName)

	for drinkName != "nothing" {
		url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + Start(drinkName)
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var response = Response{}
		json.Unmarshal(body, &response)

		fmt.Println(response.Drinks[1].Name)
		fmt.Println("----------------------------------------")
		fmt.Println("Preparation Instructions")
		str := strings.Split(response.Drinks[1].Recipe, ".")
		for _, s := range str {
			fmt.Println(s)
		}
		fmt.Println("-----------------------------------------")
		fmt.Println()

		fmt.Print("Enter drink name or enter 'nothing' to exit: ")
		fmt.Scanln(&drinkName)
	}

}
