package models

import "time"

type Shinobi struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ClanID      int       `json:"clan_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// embedded models
	Clan Clan `json:"-"`

	// open list
	// so the migrator knows that Shinobi have relationship with these three models
	NatureType []NatureType `json:"-"`
	// Jutsu      []Jutsu      `json:"-"`
}
