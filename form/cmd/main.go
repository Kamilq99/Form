package main

import (
	"database/sql"
	"log"
	"net/http"

	"form/internal/handlers"
	"form/internal/repositories"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Inicjalizacja bazy danych
	db, err := sql.Open("sqlite3", "./form.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Utworzenie tabeli
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS form_data (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL,
            email TEXT NOT NULL,
            phone TEXT NOT NULL,
            country TEXT NOT NULL,
            city TEXT NOT NULL,
            postal_code TEXT NOT NULL,
            street TEXT NOT NULL,
            house_number TEXT NOT NULL
        )
    `)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Inicjalizacja repozytorium i handlera
	repo := repositories.NewFormRepository(db)
	handler := handlers.NewFormHandler(repo)

	// Routing
	http.HandleFunc("/submit", handler.HandleFormSubmit)

	// Uruchomienie serwera
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
