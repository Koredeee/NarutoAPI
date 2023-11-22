package models

import "time"

type Jutsu struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	NatureTypeId int       `json:"nature_type_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// embedded
	NatureType NatureType `json:"-"`
}
