package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) InitRoutes() http.Handler {
	rtr := mux.NewRouter()

	// CRUD for user

	rtr.HandleFunc("/users", s.GetUsers).Methods("GET")
	rtr.HandleFunc("/user/{id:[0-9]+}", s.GetUser).Methods("GET")
	rtr.HandleFunc("/user", s.CreateUser).Methods("POST")
	rtr.HandleFunc("/user/{id:[0-9]+}", s.UpdateUser).Methods("PUT")
	rtr.HandleFunc("/user/{id:[0-9]+}", s.DeleteUser).Methods("DELETE")

	return rtr
}
