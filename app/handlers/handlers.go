package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

func RegisterHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
