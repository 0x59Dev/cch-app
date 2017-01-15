package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/methlab669/cch-app/app/handlers"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "ELO")
}

func main() {
	//Opening database connection, checking it via ping and deferring clos	e
	db, err := sql.Open("postgres", "user=testuser password=testpass dbname=testdb sslmode=disable")
	if err != nil {
		log.Println("Database connection error :", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Error with database response : ", err)
	}

	/*         stmt, err := db.Prepare("INSERT INTO users(username, email, passwd) VALUES($1, $2, $3);")
	 *
	 *         if err != nil {
	 *             log.Println("Error preparing test statement ", err)
	 *         }
	 *
	 *         res, err := stmt.Exec("DUPA", "SDSD", "SDSD")
	 *         if err != nil {
	 *             log.Println("Error executing test statement : ", err)
	 *         }
	 *         fmt.Println("------------------")
	 *         fmt.Println("DB TEST :", res) */

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
