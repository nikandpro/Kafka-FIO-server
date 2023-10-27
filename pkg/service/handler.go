package service

import (
	"encoding/json"
	"net/http"

	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

func IsCorrect(str []byte) (database.User, error) {
	user := database.User{}

	err := json.Unmarshal([]byte(str), &user)
	if err != nil {
		return user, err

	}
	return user, nil
}

func (s *Service) enrichment(user *database.User, url string) (database.User, error) {
	resp, err := http.Get(url + user.Name)
	if err != nil {
		return database.User{}, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(user)

	return *user, nil
}
