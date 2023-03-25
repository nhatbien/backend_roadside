package model

import "time"

type Order struct {
	Id           int        `json:"id,omitempty" gorm:"not null;primaryKey"`
	UserId       string     `json:"user_id,omitempty" gorm:"index"`
	RescueUnitId *string    `json:"rescue_unit_id,omitempty" gorm:"index;null"`
	Address      string     `json:"address,omitempty"`
	Lat          float64    `json:"lat,omitempty"`
	Lng          float64    `json:"lng,omitempty"`
	Note         string     `json:"note,omitempty"`
	Status       int        `json:"status,omitempty" gorm:"default:0"`
	CreatedAt    time.Time  `json:"created_at,omitempty" `
	UpdatedAt    time.Time  `json:"updated_at,omitempty" `
	User         User       `json:"user,omitempty" gorm:"foreignKey:UserId"`
	RescueUnit   RescueUnit `json:"rescue_unit,omitempty" gorm:"foreignKey:RescueUnitId"`
}
