package repositories

import (
	"database/sql"
	"errors"
	"form/internal/models"
)

type FormRepository struct {
	db *sql.DB
}

func NewFormRepository(db *sql.DB) *FormRepository {
	return &FormRepository{
		db: db,
	}
}

func (r *FormRepository) Create(form *models.Human) error {
	query := `
        INSERT INTO human (
            name, lastName, email, phone, country, city, postalCode, street, number
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query, form.Name, form.LastName, form.Email, form.Phone, form.Country, form.City, form.PostalCode, form.Street, form.Number)
	if err != nil {
		return errors.New("error creating form: " + err.Error())
	}
	return nil
}
