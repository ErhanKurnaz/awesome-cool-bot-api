package entities

type Person struct {
	BaseModel
	FirstName string    `json:"firstName" binding:"required" gorm:"type:varchar(32)"`
	LastName  string    `json:"lastName" binding:"required" gorm:"type:varchar(32)"`
	Age       int8      `json:"age" binding:"gte=1,lte=130" gorm:"type:integer"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}
