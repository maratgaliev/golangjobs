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
	ContactName    string    `json:"contact_name"`
	CurrencyType   int       `json:"currency_type"`
	EmploymentType int       `json:"employment_type"`
	Salary         *int      `json:"salary"`
	IsApproved     bool      `json:"is_approved"`
	IsRemote       bool      `json:"is_remote"`
	CreatedAt      time.Time `json:"created_at"`
}
