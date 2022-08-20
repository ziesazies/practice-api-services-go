package main

// github.com/anon-org/developing-api-services-with-golang/
// presents.clavinjune.dev/
// clavinjune.dev/fig-aug-22/
//  juneardoc@gmail.com

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const (
	queryMigrate string = `CREATE TABLE IF NOT EXISTS tasks(
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL UNIQUE,
        is_active BOOLEAN NOT NULL DEFAULT TRUE
    )`

	querySeed string = `INSERT INTO tasks (id, name) VALUES('seed1', 'task seed 1')`
)

// this function is simplified
// please don't use this function in production
func initializeDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec(queryMigrate)
	db.Exec(querySeed)
	return db
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}

		fmt.Fprintf(w, "Hello, %s!", name)
	})

	log.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)

	// You can access it at http://localhost:8000/?name=John
}
