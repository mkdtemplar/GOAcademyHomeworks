package main

import (
	"fmt"
	"sort"
	"time"
)

func sortDates(format string, dates ...string) []string {

	dateSlice := make([]string, 0)

	format = "Mar-19-2022"

	d1, _ := time.Parse(format, dates[0])
	d2, _ := time.Parse(format, dates[1])
	d3, _ := time.Parse(format, dates[2])

	date := []time.Time{d1, d2, d3}

	for df := range date {

		fmt.Println(date[df].Format(format))
	}

	sort.Slice(date, func(i, j int) bool {
		return false
	})

	for d := range date {
		dateSlice = append(dateSlice, date[d].String())
	}

	return dateSlice
}

func main() {

	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022"}

	fmt.Println(sortDates("Mar-19-2021", dates...))

}
