package server

import (
	"encoding/json"
	"net/http"
)

// Временно для теста апи
func GetAgify(w http.ResponseWriter, r *http.Request) {
	exampleAgify := `{"agify":14}`
	w.Write([]byte(exampleAgify))
}

func GetGenderize(w http.ResponseWriter, r *http.Request) {
	exampleAgify := `{"genderize":"M"}`
	w.Write([]byte(exampleAgify))
}

func GetNationalize(w http.ResponseWriter, r *http.Request) {
	exampleAgify := `{"nationalize":"Russia"}`
	w.Write([]byte(exampleAgify))
}


// controllers for rest api
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := `[
			{
			"name":"John"
			},
			{
				"name":"John"
			}, 
			]`
	json_data, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(json_data)

}

func GetUser() {}

func CreateUser() {}

func UpdateUser() {}

func DeleteUser() {}
