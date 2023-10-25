package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

func IsCorrect(str []byte) (database.User, error) {
	user := database.User{}

	err := json.Unmarshal([]byte(str), &user)
	if err != nil {
		log.Fatal(err)
		return user, err

	}
	return user, nil
}

func (s *Service) enrichment(user *database.User) (database.User, error) {
	resp, err := http.Get("http://localhost:8081/&name=" + user.Name)
	if err != nil {
		log.Fatal("Bad request", err)
	}

	defer resp.Body.Close()

	// json.Decoder(resp.Body).Decode(&result)

	return *user, nil
}
