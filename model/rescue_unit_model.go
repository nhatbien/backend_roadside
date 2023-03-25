package model

import "time"

type RescueUnit struct {
	Id        string    `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty"`
	Phone     string    `json:"phone,omitempty" gorm:"size:255;uniqueIndex"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty" gorm:"size:255;uniqueIndex"`
	Address   string    `json:"address,omitempty"`
	Lat       float64   `json:"lat,omitempty"`
	Lng       float64   `json:"lng,omitempty"`
	Type      string    `json:"type,omitempty"`
	Status    int       `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Token     string    `json:"token,omitempty"  gorm:"-"`
}
