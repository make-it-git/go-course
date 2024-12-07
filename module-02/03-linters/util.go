package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func unsafeUtil(db *sql.DB) {
	username := "admin' OR '1'='1"
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)
	fmt.Println("Unsafe Query:", query)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	os.Chmod("/tmp/123", os.FileMode(0644))
}
