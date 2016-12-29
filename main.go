package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/methlab669/cch-app/app/handlers"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "ELO")
}
func main() {

	logFile, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Cannot create/open logfile : ", err)
	}
	defer logFile.Close()
	//Router section of main server file

	router := httprouter.New()

	router.POST("/register", handlers.Register)

	log.Println("Server is up and listening")
	log.Fatal(http.ListenAndServe(":9000", router))
}
