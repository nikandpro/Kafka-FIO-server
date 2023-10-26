package database

type Database interface {
	Get() error
	Create(User) error
	// Delete(id string)
	// Update(id string)
}

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Agify       int    `json:"agify"`
	Genderize   string `json:"genderize"`
	Nationalize string `json:"nationalize"`
}
