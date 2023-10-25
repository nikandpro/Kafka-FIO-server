package postgres

import (
	"database/sql"
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
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Patronymic)
		if err != nil {
			log.Fatal("error get next", err)
			continue
		}
		users = append(users, u)
	}

	// for _, u := range users {
	// 	fmt.Println(u.ID, u.Name, u.Surname, u.Patronymic)
	// }

	return nil
}

func (db *PostgresDB) Create(database.User) error {
	// insert, err := db.connection.Query("INSERT INTO articles (title, anons, full_text) VALUES ( ?, ?, ?)", post.Anons, post.Anons, post.FullText)
	// if err != nil {
	// 	log.Fatal("error Create ", err)
	// 	return err
	// }
	// defer insert.Close()

	return nil
}
