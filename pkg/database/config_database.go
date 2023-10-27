package database

type Database interface {
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(User) error
	UpdateUser(user User, id string) error
	DeleteUser(id string) error
}

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Country    string `json:"country"`
}
