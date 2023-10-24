package database

type Database interface {
	Get() error
	Create([]byte) error
	IsCorrect(str []byte) (User, error)
	// IsCorrect(str []byte) (User, error)
	// Delete(id string)
	// Update(id string)
}

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
