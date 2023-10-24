package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nikandpro/kafka-fio-server/pkg/config"
	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

func init() {
	connStr := ""
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("create table users(id serial primary key, name varchar(50), surname varchar(50), patronymic varchar(50));")

}

type PostgresDB struct {
	connection *sql.DB
}

func New(cfg *config.Config) (*PostgresDB, error) {
	connStr := cfg.DBPath
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return &PostgresDB{}, err
	}
	return &PostgresDB{
		connection: conn,
	}, nil
}

func (db *PostgresDB) Get() error {
	rows, err := db.connection.Query("select * from users")
	if err != nil {
		fmt.Println("error Get")
		return err
	}
	defer rows.Close()
	users := []database.User{}

	for rows.Next() {
		u := database.User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic)
		if err != nil {
			fmt.Println("error get next", err)
			continue
		}
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Surname, u.Patronymic)
	}

	return nil
}

func (db *PostgresDB) Create(json []byte) error {
	fmt.Println("create:")
	fmt.Println(string(json))
	return nil
}

func (db *PostgresDB) IsCorrect(str []byte) (database.User, error) {
	user := database.User{}

	err := json.Unmarshal([]byte(str), &user)
	if err != nil {
		log.Fatal(err)
		return user, err

	}
	return user, nil
}
