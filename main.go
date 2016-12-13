package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/methlab669/cavechat/app/utils/router"
)

func main() {

	logFile, err := os.OpenFile("SampleLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Cannot create/open logfile : ", err)
	}
	defer logFile.Close()

	URLrouter := router.NewRouter()
	log.Fatal(http.ListenAndServe(":9000", URLrouter))
	fmt.Println("Server listening at localhost:8080")
}
