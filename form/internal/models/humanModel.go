package models

type Human struct {
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Country    string `json:"country"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Street     string `json:"street"`
	Number     string `json:"number"`
}
