package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

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
