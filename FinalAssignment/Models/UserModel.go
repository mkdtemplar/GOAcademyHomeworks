package Models

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" `
	Password string `json:"password"`
}

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
