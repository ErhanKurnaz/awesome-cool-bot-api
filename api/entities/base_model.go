package entities

import (
	"time"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IBaseModel interface {
	SetId(uuid.UUID)
}

type BaseModel struct {
	ID        uuid.UUID  `json:"id"        gorm:"primary_key;type:uuid"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	DeleteAt  *time.Time `json:"deletedAt"                                                sql:"index"`
}

func (entity *BaseModel) SetId(id uuid.UUID) {
	entity.ID = id
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.NewV4()

	base.ID = uuid
	return nil
}
