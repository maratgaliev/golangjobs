package model

import "time"

// Job entity, represents job
// swagger:model
type Job struct {
	Id             int       `json:"id"`
	Title          string    `json:"title" validate:"required"`
	Description    string    `json:"description" validate:"required"`
	City           string    `json:"city" validate:"required"`
	Email          string    `json:"email" validate:"required"`
	Company        string    `json:"company" validate:"required"`
	Phone          string    `json:"phone"`
	Website        string    `json:"website" validate:"required"`
	ContactName    string    `json:"contact_name"`
	CurrencyType   string    `json:"currency_type"`
	EmploymentType string    `json:"employment_type"`
	Salary         *int      `json:"salary"`
	IsApproved     bool      `json:"is_approved"`
	CreatedAt      time.Time `json:"created_at"`
}
