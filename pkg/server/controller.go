package server

import (
	"encoding/json"
	"net/http"
)

// Временно для теста апи
func GetAgify(w http.ResponseWriter, r *http.Request) {
	exampleAgify := `"agify":14`
	json_data, err := json.Marshal(exampleAgify)
	if err != nil {
		panic(err)
	}
	w.Write(json_data)
}

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
