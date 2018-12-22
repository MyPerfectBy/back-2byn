package model

import "time"

type Product struct {
	ID int     `gorm:"AUTO_INCREMENT" json:"id"`
	Title string `json:"title"`
	PhotoURL string `json:"photoUrl"`
	Description string `json:"description"`
	Contacts string `json:"contacts"`
	Approved bool `json:"approved"`
	Date time.Time `json:"date"`
}
