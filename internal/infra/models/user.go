package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	User      string    `gorm:"unique"`
	Name      string    `conform:"name,trim"`
	Surname   string    `conform:"trim"`
	CreatedAt time.Time `sql:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" gorm:"column:createdAt"`
	UpdatedAt time.Time `sql:"default:CURRENT_TIMESTAMP" gorm:"column:updatedAt"`
}
