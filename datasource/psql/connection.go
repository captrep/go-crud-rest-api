package psql

import (
	"database/sql"
	"log"

	"github.com/captrep/go-crud-rest-api/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() *sql.DB {
	db, err := sql.Open(config.Conf.DBDriver, config.Conf.DBSource)
	if err != nil {
		log.Println(err)
	}
	errPing := db.Ping()
	if errPing != nil {
		panic(errPing)
	}

	log.Println("Database succesfully configured!")
	return db
}
