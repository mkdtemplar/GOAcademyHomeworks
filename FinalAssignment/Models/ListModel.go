package Models

type Lists struct {
	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type CreateLists struct {
	Name string `json:"name" binding:"required"`
}
