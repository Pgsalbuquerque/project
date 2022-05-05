package dto

import (
	"time"
)

type ProjectResponseDTO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Users       []string  `json:"users"`
	Admins      []string  `json:"admins"`
	Owner       string    `json:"owner"`
	Nda_accept  bool      `json:"nda_accept"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
}
