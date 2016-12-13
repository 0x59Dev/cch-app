package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/methlab669/cavechat/app/models/auth"
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
	fmt.Println(registerData)

}
