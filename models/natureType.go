package models

import "time"

type NatureType struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// put the ShinobiID so it can GET NatureType by ShinobiID
	// ShinobiID will called by shinobiController
	ShinobiID int       `json:"shinobi_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// embedded the model
	Shinobi Shinobi `json:"-"`

	Jutsu []Jutsu `json:"-"`
}
