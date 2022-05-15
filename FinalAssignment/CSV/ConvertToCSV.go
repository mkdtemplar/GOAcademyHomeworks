package CSV

import (
	models "FinalAssignment/Models"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	for i := 0; i < len(list); i++ {
		var row []string
		row = append(row, strconv.Itoa(int(list[i].Id)))
		row = append(row, list[i].Name)
		_ = csvWriter.Write(row)
	}
	csvWriter.Flush()
	csvFile.Close()
	c.JSON(http.StatusOK, list)
}
