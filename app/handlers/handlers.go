package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type registerForm struct {
	Login        string `json:"login"`
	Email        string `json:"mail"`
	Password     string `json:"password"`
	ConfPassword string `json:"confPassword"`
}

type TestForm struct {
	Test  string `json:"test"`
	Test2 string `json:"test2"`
}

func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	var registerJSON registerForm
	JSONDecoder := json.NewDecoder(r.Body)
	err := JSONDecoder.Decode(&registerJSON)
	fmt.Println(registerJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error occured while encoding test json : ", err)
		return
	}

}
