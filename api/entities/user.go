package entities

type User struct {
	BaseModel
	Email string `json:"email" binding:"required,email" gorm:"type:varchar(256);UNIQUE"`
	Name  string `json: "name" binding:"required,min=3,max=256" gorm:"type:varchar(256)"`
}
