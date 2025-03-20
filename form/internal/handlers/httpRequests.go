package handlers

import (
	"fmt"
	"net/http"

	"form/internal/models"
	"form/internal/repositories"
)

type FormHandler struct {
	repo *repositories.FormRepository
}

func NewFormHandler(repo *repositories.FormRepository) *FormHandler {
	return &FormHandler{repo: repo}
}

func (h *FormHandler) HandleFormSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parsowanie formularza
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Odczytanie danych z formularza
	formData := models.Human{
		Name:       r.FormValue("first_name"),
		LastName:   r.FormValue("last_name"),
		Email:      r.FormValue("email"),
		Phone:      r.FormValue("phone"),
		Country:    r.FormValue("country"),
		City:       r.FormValue("city"),
		PostalCode: r.FormValue("postal_code"),
		Street:     r.FormValue("street"),
		Number:     r.FormValue("house_number"),
	}

	// Zapisanie danych do bazy danych
	err = h.repo.Create(&formData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to save data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Form submitted successfully!")
}
