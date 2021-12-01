package database

import (
	"database/sql"
	"fmt"
	"kasir/config"
	"kasir/database/postgres"
	"log"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	Conn    *sql.DB
	Queries *postgres.Queries
}

func NewPostgres(config config.Config) (postgreSQL PostgreSQL, err error) {
	connInf := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Db_host,
		config.Db_port,
		config.Db_username,
		config.Db_password,
		config.Db_name,
		config.Db_sslmode,
	)

	sqlconn, err := sql.Open(config.Db_driver, connInf)
	if err != nil {
		log.Fatal(err)
		return
	}

	postgreSQL.Conn = sqlconn
	postgreSQL.Queries = postgres.New(sqlconn)
	return
}
