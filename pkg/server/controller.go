package server

import (
	"encoding/json"
	"net/http"
)

// Временно для теста апи
func GetAgify(w http.ResponseWriter, r *http.Request) {
	exampleAgify := `{
		"agify": 12
	}`
	w.Write([]byte(exampleAgify))
}

func GetGenderize(w http.ResponseWriter, r *http.Request) {
	exampleGender := `{
		"genderize": "Male"
	}`
	w.Write([]byte(exampleGender))
}

func GetNationalize(w http.ResponseWriter, r *http.Request) {
	exampleNation := `{
		"nationalize": "Russion"
	}`
	w.Write([]byte(exampleNation))
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
