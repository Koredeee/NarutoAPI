package models

import "time"

type Shinobi struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// put the ClanID so it can GET Shinobi by Clan ID
	ClanID    int       `json:"clan_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// embedded models
	Clan Clan `json:"-"`

	// open list
	// so the migrator knows that Shinobi have relationship with NatureType
	NatureType []NatureType `json:"-"`
	// Jutsu      []Jutsu      `json:"-"`
}
