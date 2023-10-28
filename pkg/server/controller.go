package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

// Временно для теста апи
func GetAgify(w http.ResponseWriter, r *http.Request) {
	exampleAgify := `{"count":3800,"name":"Dmitriy","age":42}`
	w.Write([]byte(exampleAgify))
}

func GetGenderize(w http.ResponseWriter, r *http.Request) {
	exampleGender := `{"count":25459,"name":"Dmitriy","gender":"male","probability":1.0}`
	w.Write([]byte(exampleGender))
}

func GetNationalize(w http.ResponseWriter, r *http.Request) {
	exampleNation := `{"count":24968,"name":"Dmitriy","country":[{"country_id":"UA","probability":0.419},{"country_id":"RU","probability":0.291},{"country_id":"KZ","probability":0.097},{"country_id":"BY","probability":0.069},{"country_id":"IL","probability":0.019}]}`
	w.Write([]byte(exampleNation))
}

// controllers for rest api
func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.db.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	json_data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json_data)

}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	user, err := s.db.GetUser(id["id"])
	if err != nil {
		log.Fatal(err)
	}
	json_data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json_data)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("update api")
	var user = database.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	var user = database.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.UpdateUser(user, id["id"])
	if err != nil {
		log.Fatal(err)
	}

}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	err := s.db.DeleteUser(id["id"])
	if err != nil {
		log.Fatal(err)
	}
}
