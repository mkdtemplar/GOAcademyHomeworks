package Controllers

import (
	model "FinalAssignment/Repository/Models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetWeather(c *gin.Context) {

	weatherH := strings.SplitN(c.Request.Header.Get("Weather"), " ", 2)

	payload, _ := base64.StdEncoding.DecodeString(weatherH[0])
	pair := strings.SplitN(string(payload), ":", 2)

	lat, _ := strconv.ParseFloat(pair[0], 64)
	lon, _ := strconv.ParseFloat(pair[1], 64)

	url := fmt.Sprintf("%s%f%s%f%s%s", "https://api.openweathermap.org/data/2.5/weather?lat=", lat, "&lonHeader=",
		lon, "&appid=", "1c3d10c3307ce3d7f22757f9fbf51020")

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var response = model.Weather{}
	json.Unmarshal(body, &response)

	var response1 = model.Response{}

	response1 = model.Response{
		Description: response.Weather[0].Description,
		Temperature: fmt.Sprintf("%.2f%s", response.Main.Temp-273.15, " Celsius"),
		City:        response.Name,
	}

	c.JSON(http.StatusOK, response1)

}
