package postgres

import (
	"database/sql"
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

	db.Exec("create table users(id serial primary key, name varchar(50), surname varchar(50), patronymic varchar(50), agify int, genderize varchar(50), nationalize varchar(50));")

}

type PostgresDB struct {
	connection *sql.DB
}

func New(cfg *config.Config) (*PostgresDB, error) {
	conn, err := sql.Open("postgres", cfg.DBPath)
	if err != nil {
		log.Fatal("error New DB", err)
		return &PostgresDB{}, err
	}
	return &PostgresDB{
		connection: conn,
	}, nil
}

func (db *PostgresDB) Get() error {
	rows, err := db.connection.Query("select * from users")
	if err != nil {
		log.Fatal("error Get", err)
		return err
	}
	defer rows.Close()
	users := []database.User{}

	for rows.Next() {
		u := database.User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic, &u.Agify, &u.Genderize, &u.Nationalize)
		if err != nil {
			log.Fatal("error get next", err)
			continue
		}
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Surname, u.Patronymic, u.Agify, u.Genderize, u.Nationalize)
	}

	return nil
}

func (db *PostgresDB) Create(user database.User) error {
	psgQuery := fmt.Sprintf(`insert into users(name, surname, patronymic, agify, genderize, nationalize) values ('%s', '%s', '%s', %d, '%s', '%s')`, user.Name, user.Surname, user.Patronymic, user.Agify, user.Genderize, user.Nationalize)
	fmt.Println(psgQuery)
	insert, err := db.connection.Query(psgQuery)
	if err != nil {
		log.Fatal("error Create ", err)
		return err
	}
	defer insert.Close()

	return nil
}
