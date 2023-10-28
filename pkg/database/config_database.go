package database

type Database interface {
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(User) error
	UpdateUser(user User, id string) error
	DeleteUser(id string) error
	CreateCountry(country Country) error
	CreateUser_Country(user_id int, country_id string) error
}

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	Country    []Country `json:"country"`
}

type Country struct {
	ID          string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
