package models

import (
	"time"
)

type Actor struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Salary    float64   `json:"salary"`
}
