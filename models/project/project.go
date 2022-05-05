package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Users       pq.StringArray `gorm:"type:text[]" json:"users"`
	Admins      pq.StringArray `gorm:"type:text[]" json:"admins"`
	Owner       string         `json:"owner"`
	Nda_accept  bool           `json:"nda_accept"`
	Link        string         `json:"link"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}
