package helper

import (
	"log"
	"os"
)

func EnsureDBDirectory() {
	const dbDir = "db"
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err := os.Mkdir(dbDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory %s: %v", dbDir, err)
		}
		log.Printf("Directory %s created successfully", dbDir)
	}
}
