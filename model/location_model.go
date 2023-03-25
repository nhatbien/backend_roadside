package model

type Location struct {
	Id      int     `json:"id,omitempty" gorm:"primaryKey"`
	Lat     float64 `json:"lat,omitempty" gorm:"not null"`
	Lng     float64 `json:"lng,omitempty"  gorm:"not null"`
	Address string  `json:"address,omitempty"`
}
