package main

import (
	"log"
	"net/http"
	"os"

	"github.com/methlab669/cch-app/app/utils/router"
)

func main() {

	logFile, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Cannot create/open logfile : ", err)
	}
	defer logFile.Close()

	URLrouter := router.NewRouter()
	log.Println("Server listeninadsasdg at localhost:8080")
	log.Fatal(http.ListenAndServe(":9000", URLrouter))
}
