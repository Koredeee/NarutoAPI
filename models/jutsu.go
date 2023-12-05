package models

import "time"

type Jutsu struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`

	// put the NatureTypeID so it can GET Jutsu by NatureType ID
	NatureTypeID int       `json:"nature_type_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// embedded
	NatureType NatureType `json:"-"`
}
