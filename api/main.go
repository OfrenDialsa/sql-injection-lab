package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func Run() {
	var err error

	db, err = sql.Open("sqlite", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	initDB()

	http.HandleFunc("/user", userRoute)
	http.HandleFunc("/login", loginRoute)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);

	DELETE FROM users;

	INSERT INTO users (username, password) VALUES
	('admin', 'admin123'),
	('nero', 'secret'),
	('guest', 'guest');
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// you can inject sql query inside this route
func userRoute(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	query := fmt.Sprintf("SELECT id, username FROM users WHERE id = %s", id)

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		rows.Scan(&id, &username)
		fmt.Fprintf(w, "ID: %d | Username: %s\n", id, username)
	}
}

// this one for auth bypass
func loginRoute(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	query := fmt.Sprintf(
		"SELECT id FROM users WHERE username = '%s' AND password = '%s'",
		username, password,
	)

	row := db.QueryRow(query)

	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Fprintln(w, "Login failed")
		return
	}

	fmt.Fprintf(w, "Login success! User ID: %d\n", id)
}
