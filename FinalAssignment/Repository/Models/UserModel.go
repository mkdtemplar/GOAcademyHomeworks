package Models

type User struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
