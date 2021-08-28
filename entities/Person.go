package entities

type Person struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}
