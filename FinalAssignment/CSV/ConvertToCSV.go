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
	var task []models.Tasks
	models.DB.Find(&list)
	models.DB.Find(&task)
	_ = csvWriter.Write([]string{"Id", "Name"})
	for i := 0; i < len(list); i++ {
		var row []string
		row = append(row, strconv.Itoa(int(list[i].Id)))
		row = append(row, list[i].Name)
		_ = csvWriter.Write(row)
	}

	_ = csvWriter.Write([]string{"id", "Text", "list_id", "completed"})

	for i := 0; i < len(task); i++ {
		var taskRow []string
		taskRow = append(taskRow, strconv.Itoa(int(task[i].Id)))
		taskRow = append(taskRow, task[i].Text)
		taskRow = append(taskRow, strconv.Itoa(task[i].ListId))
		taskRow = append(taskRow, strconv.FormatBool(task[i].Completed))
		_ = csvWriter.Write(taskRow)
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
