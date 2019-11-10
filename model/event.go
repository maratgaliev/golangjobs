package model

import "time"

// Event entity, represents event
// swagger:model

type Event struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Company     string    `json:"company" validate:"required"`
	IsApproved  bool      `json:"is_approved"`
	CreatedAt   time.Time `json:"created_at"`
}
