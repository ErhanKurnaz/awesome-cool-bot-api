package entities

type Video struct {
	ID          uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Title       string `json:"title" binding:"required,min=2,max=100" validate:"is-cool" gorm:"type:varchar(100)"`
	Description string `json:"description" binding:"required,max=200" gorm:"type:varchar(200)"`
	URL         string `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person `json:"author" binding:"required" gorm:"foreignkey:AuthorID"`
	AuthorID	uint64 `json:"-"`
}
