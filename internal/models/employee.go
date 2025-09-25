package models

import (
	"time"
)

type Employee struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	Email      string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Position   string    `gorm:"size:100" json:"position"`
	Department string    `gorm:"size:100" json:"department"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
