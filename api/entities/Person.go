package entities

import "time"

type Person struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string    `json:"firstName" binding:"required" gorm:"type:varchar(32)"`
	LastName  string    `json:"lastName" binding:"required" gorm:"type:varchar(32)"`
	Age       int8      `json:"age" binding:"gte=1,lte=130" gorm:"type:integer"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}
