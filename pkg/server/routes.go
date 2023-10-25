package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() http.Handler {
	rtr := mux.NewRouter()

	//testing enrich api

	rtr.HandleFunc("/agify/{name}", GetAgify).Methods("GET")
	rtr.HandleFunc("/genderize/{name}", GetGenderize).Methods("GET")
	rtr.HandleFunc("/nationalize/{name}", GetNationalize).Methods("GET")


	rtr.HandleFunc("/users", GetUsers).Methods("GET")
	// rtr.HandleFunc("/user/{id:[0-9]+}", GetUser).Methods("GET")
	// rtr.HandleFunc("/user/", CreateUser).Methods("POST")
	// rtr.HandleFunc("/user/{id:[0-9]+}", UpdateUser).Methods("UPDATE")
	// rtr.HandleFunc("/user/{id:[0-9]+}", DeleteUser).Methods("DELETE")

	// //authentication
	// rtr.HandleFunc("/login", handlers.Login).Methods("POST")

	// //CRUD User
	// rtr.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	// rtr.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	// rtr.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	// rtr.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	// rtr.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// //CRUD Posts
	// //CRUD User
	// rtr.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	// rtr.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
	// rtr.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	// rtr.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
	// rtr.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

	return rtr
}
