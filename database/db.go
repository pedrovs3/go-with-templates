package database

import (
	u "alura_store/utils"
	"database/sql"
)

func ConnectWithDatabase() *sql.DB {
	connStr := "user=postgres dbname=alura_store password=Ccx967408 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	u.Check(err)

	return db
}
