package db

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInsertUser(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %s", err)
	}
	defer db.Close()

	createTableQuery := `
		CREATE TABLE USERS (
			user_id TEXT PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("failed to create table: %s", err)
	}

	userModel := &UserModel{DB: db}

	// Define test inputs
	id := "abcd"
	username := "aaochieng"
	email := "aaochieng@example.com"
	password := "securepassword"

	err = userModel.InsertUser(id, username, email, password)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	var (
		gotID       string
		gotUsername string
		gotEmail    string
		gotPassword string
	)

	query := "SELECT user_id, username, email, password FROM USERS WHERE user_id = ?"
	row := db.QueryRow(query, id)
	err = row.Scan(&gotID, &gotUsername, &gotEmail, &gotPassword)
	if err != nil {
		t.Fatalf("failed to query inserted user: %s", err)
	}

	if gotID != id || gotUsername != username || gotEmail != email || gotPassword != password {
		t.Errorf("inserted user data does not match: got (%s, %s, %s, %s), want (%s, %s, %s, %s)",
			gotID, gotUsername, gotEmail, gotPassword, id, username, email, password)
	}
}



func TestUserExists(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %s", err)
	}
	defer db.Close()

	createTableQuery := `
		CREATE TABLE USERS (
			user_id TEXT PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("failed to create table: %s", err)
	}

	userModel := &UserModel{DB: db}

	insertUserQuery := `
		INSERT INTO USERS (user_id, username, email, password)
		VALUES ('1', 'squidward', 'squidward@example.com', 'securepassword');`
	_, err = db.Exec(insertUserQuery)
	if err != nil {
		t.Fatalf("failed to insert test user: %s", err)
	}

	tests := []struct {
		email       string
		expected    bool
		shouldError bool
	}{
		{"squidward@example.com", true, false},  // User exists
		{"nonexistent@example.com", false, false}, // User does not exist
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			// Call the method under test
			exists, err := userModel.UserExists(test.email)

			// Check for unexpected errors
			if test.shouldError && err == nil {
				t.Errorf("expected an error but got none")
			}
			if !test.shouldError && err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			// Verify the result
			if exists != test.expected {
				t.Errorf("unexpected result: got %v, want %v", exists, test.expected)
			}
		})
	}
}
