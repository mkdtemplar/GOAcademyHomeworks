package main

import (
	models "FinalAssignment/Repository/DatabaseContext"
	routers "FinalAssignment/RoutersSetup"
	"FinalAssignment/cmd"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	models.ConnectDatabase()

	router := routers.Setup()

	exit := make(chan struct{}, 1)
	// start the server
	go func() {
		log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
		exit <- struct{}{}
	}()

	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:3000/swagger").Start()
	case "darwin":
		err = exec.Command("open", "http://localhost:3000/swagger").Start()
	case "linux":
		err = exec.Command("xdg-open", "http://localhost:3000/swagger").Start()
	}
	if err != nil {
		fmt.Println(err)
	}
	<-exit

}
