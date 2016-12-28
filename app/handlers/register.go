package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/methlab669/cch-app/app/models/auth"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var registerData auth.RegisterData

	JSONDecoder := json.NewDecoder(r.Body)

	err := JSONDecoder.Decode(&registerData)
	if err != nil {
		errorString := "ooops! We got some problem, please try again!"
		log.Println("Error Encoding registration data : ", err)
		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}
	if err := registerData.AddUser(); err != nil {
		log.Println("Error : ", err)
		return
	} else {
		log.Println("Success")
	}
}
