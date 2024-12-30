package helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const sqlFilePath = "sql/user_table.sql"

func CreateTableFromSQL(db *sql.DB) error {
	tableExists, err := checkIfTableExists(db, "users")

	if err != nil {
		log.Printf("Failed to check if table exists: %v\n", err)
		return err
	}

	if tableExists {
		log.Println("Users table already exists. Skipping creation.")
		return nil
	}

	sqlContent, err := os.ReadFile(sqlFilePath)
	if err != nil {
		log.Printf("Failed to read SQL file: %v\n", err)
		return err
	}

	_, err = db.Exec(string(sqlContent))
	if err != nil {
		log.Printf("Failed to execute SQL: %v\n", err)
		return err
	}

	log.Println("Users table is ready.")
	return nil
}

func checkIfTableExists(db *sql.DB, tableName string) (bool, error) {
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", tableName)
	var name string
	err := db.QueryRow(query).Scan(&name)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error while checking for table %s: %v\n", tableName, err)
		return false, err
	}

	return name == tableName, nil
}