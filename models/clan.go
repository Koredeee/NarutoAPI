package models

import "time"

type Clan struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// open the list
	// just let the migrator knows that Clan has relationship with Shinobi
	Shinobi []Shinobi `json:"-"`
}
