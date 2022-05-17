package CSV

import (
	models "FinalAssignment/Models"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

func ReadListRow(c *gin.Context) {
	csvFile, err := os.Create("CSV/CSVFileOutput/list.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	var list []models.Lists
	models.DB.Find(&list)
	_ = csvWriter.Write([]string{"Id", "Name"})
	for i := 0; i < len(list); i++ {
		var row []string
		row = append(row, strconv.Itoa(int(list[i].Id)))
		row = append(row, list[i].Name)
		_ = csvWriter.Write(row)
	}
	if err != nil {
		log.Fatal(err)
	}
	fileName := "list.csv"
	path := "CSV/CSVFileOutput/list.csv"

	csvWriter.Flush()
	csvFile.Close()
	c.FileAttachment(path, fileName)
}
