package main

import (
	"fmt"
	"sort"
	"time"
)

type timeSlice []time.Time

func (t timeSlice) Len() int {
	return len(t)
}

func (t timeSlice) Less(i, j int) bool {
	return t[i].Before(t[j])
}

func (t timeSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func sortDates(format string, dates ...string) []string {

	dateSlice := make([]string, 0)
	date1 := dates[0]
	date2 := dates[1]
	date3 := dates[2]
	date4 := dates[3]

	d1, _ := time.Parse(format, date1)
	d2, _ := time.Parse(format, date2)
	d3, _ := time.Parse(format, date3)
	d4, _ := time.Parse(format, date4)

	var date timeSlice = []time.Time{d1, d2, d3, d4}

	sort.Sort(date)

	for d := range date {
		dateSlice = append(dateSlice, date[d].Format(format))
	}

	return dateSlice
}

func main() {

	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022", "Dec-16-1975"}
	const format = "Jan-02-2006"

	fmt.Println(sortDates(format, dates...))

}
