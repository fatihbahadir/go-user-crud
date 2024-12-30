package config

import (
	"database/sql"
	"log"
	"user-crud/helper"
)

type DBPath string

const (
	dbPath DBPath = "db/test.db"
)

func DatabaseConnection() *sql.DB {
	helper.EnsureDBDirectory()

	db, err := sql.Open("sqlite3", string(dbPath))
	helper.HandleError(err, "Failed to connect to the database")


	err = db.Ping()
	helper.HandleError(err, "Failed to verify the database connection")

	log.Println("Connected to Database")
	
	err = helper.CreateTableFromSQL(db)
	helper.HandleError(err, "Failed to create the table")

	return db;
}