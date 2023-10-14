package database

type Database interface {
	Get() error
	Create([]byte) error
	// Delete(id string)
	// Update(id string)
}
