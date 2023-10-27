package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nikandpro/kafka-fio-server/pkg/config"
	"github.com/nikandpro/kafka-fio-server/pkg/database"
)

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

func (db *PostgresDB) GetUsers() ([]database.User, error) {
	rows, err := db.connection.Query("select * from users order by id")
	if err != nil {
		log.Fatal("error get", err)
		return nil, err
	}
	defer rows.Close()
	users := []database.User{}

	for rows.Next() {
		u := database.User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic, &u.Agify, &u.Genderize, &u.Nationalize)
		if err != nil {
			log.Fatal("error get users", err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func (db *PostgresDB) GetUser(id string) (database.User, error) {
	rows, err := db.connection.Query("select * from users where id = " + id)
	if err != nil {
		log.Fatal("error get id=", id, err)
		return database.User{}, err
	}
	defer rows.Close()
	u := database.User{}

	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic, &u.Agify, &u.Genderize, &u.Nationalize)
		if err != nil {
			log.Fatal("error get user", err)
			continue
		}
	}

	return u, nil
}

func (db *PostgresDB) CreateUser(user database.User) error {
	psgQuery := fmt.Sprintf(`insert into users(name, surname, patronymic, agify, genderize, nationalize) values ('%s', '%s', '%s', %d, '%s', '%s')`, user.Name, user.Surname, user.Patronymic, user.Agify, user.Genderize, user.Nationalize)
	insert, err := db.connection.Query(psgQuery)
	if err != nil {
		log.Fatal("error Create user", err)
		return err
	}
	defer insert.Close()

	return nil
}

func (db *PostgresDB) UpdateUser(user database.User, id string) error {
	psgQuery := fmt.Sprintf(`update users
	set name='%s',
	surname='%s',
	patronymic='%s',
	agify=%d,
	genderize='%s',
	nationalize='%s'
	where id = %s`, user.Name, user.Surname, user.Patronymic, user.Agify, user.Genderize, user.Nationalize, id)
	update, err := db.connection.Query(psgQuery)
	if err != nil {
		log.Fatal("error Create user", err)
		return err
	}
	defer update.Close()

	return nil
}

func (db *PostgresDB) DeleteUser(id string) error {
	psgQuery := fmt.Sprintf(`delete from users where id = %s`, id)
	_, err := db.connection.Query(psgQuery)
	if err != nil {
		log.Fatal("error Create user", err)
		return err
	}

	return nil
}
