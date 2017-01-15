package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/methlab669/cch-app/app/models/auth"
)

var db *sql.DB

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var registerData auth.User

	JSONDecoder := json.NewDecoder(r.Body)

	err := JSONDecoder.Decode(&registerData)
	if err != nil {
		log.Println("Error Encoding registration data : ", err)
		status := http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
		return
	}

	fmt.Println(registerData)
	if err := registerData.AddUser(); err != nil {
		log.Println("Error : ", err)
		return
	} else {
		log.Println("Success")
	}
}
