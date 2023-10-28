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
		return &PostgresDB{}, err
	}
	return &PostgresDB{
		connection: conn,
	}, nil
}

func (db *PostgresDB) GetUsers() ([]database.User, error) {
	rows, err := db.connection.Query("select * from users order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []database.User{}

	for rows.Next() {
		u := database.User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic, &u.Age, &u.Gender)
		if err != nil {
			log.Fatal(err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func (db *PostgresDB) GetUser(id string) (database.User, error) {
	rows, err := db.connection.Query("select * from users where id = " + id)
	if err != nil {
		return database.User{}, err
	}
	defer rows.Close()
	u := database.User{}

	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic, &u.Age, &u.Gender)
		if err != nil {
			log.Fatal(err)
			continue
		}
	}

	return u, nil
}

func (db *PostgresDB) CreateUser(user database.User) error {
	psgQuery := fmt.Sprintf(`insert into users(name, surname, patronymic, age, gender) values ('%s', '%s', '%s', '%d', '%s') returning id`, user.Name, user.Surname, user.Patronymic, user.Age, user.Gender)
	insert, err := db.connection.Prepare(psgQuery)
	if err != nil {
		return err
	}

	err = insert.QueryRow().Scan(&user.ID)
	if err != nil {
		return err
	}

	defer insert.Close()
	for _, c := range user.Country {
		db.CreateCountry(c)
		err = db.CreateUser_Country(user.ID, c.ID)
		if err != nil {
			log.Fatal("don", err)
			return err
		}

	}

	return nil
}

func (db *PostgresDB) UpdateUser(user database.User, id string) error {
	psgQuery := fmt.Sprintf(`update users
	set name='%s',
	surname='%s',
	patronymic='%s',
	age=%d,
	gender='%s'
	where id = %s`, user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, id)
	update, err := db.connection.Query(psgQuery)
	if err != nil {
		return err
	}
	defer update.Close()

	return nil
}

func (db *PostgresDB) DeleteUser(id string) error {
	psgQuery := fmt.Sprintf(`delete from users where id = %s`, id)
	_, err := db.connection.Query(psgQuery)
	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresDB) CreateCountry(country database.Country) error {
	psgQuery := fmt.Sprintf(`insert into country(id, probability) values ('%s', '%f')`, country.ID, country.Probability)
	insert, err := db.connection.Query(psgQuery)
	if err != nil {
		return err
	}

	defer insert.Close()

	return nil
}

func (db *PostgresDB) CreateUser_Country(user_id int, country_id string) error {
	psgQuery := fmt.Sprintf(`insert into users_country(user_id, country_id) values ('%d', '%s')`, user_id, country_id)
	insert, err := db.connection.Query(psgQuery)
	if err != nil {
		return err
	}

	defer insert.Close()

	return nil
}
