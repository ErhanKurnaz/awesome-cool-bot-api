package entities

import "time"

type BaseModel struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}
