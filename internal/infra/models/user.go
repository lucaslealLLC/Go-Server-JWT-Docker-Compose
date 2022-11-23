package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	User      string    `json:"user" gorm:"unique"`
	Name      string    `json:"name" conform:"name,trim"`
	Surname   string    `json:"surname" conform:"trim"`
	CreatedAt time.Time `json:"createdAt" sql:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" sql:"default:CURRENT_TIMESTAMP" gorm:"column:updatedAt"`
}
