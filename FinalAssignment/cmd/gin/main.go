package main

import (
	models "FinalAssignment/Repository/DatabaseContext"
	routers "FinalAssignment/RoutersSetup"
	"FinalAssignment/cmd"
	"log"
	"net/http"
)

func main() {

	models.ConnectDatabase()

	router := routers.Setup()

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))

}
