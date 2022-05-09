package Models

type Task struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Text      string `json:"text"`
	ListId    uint   `json:"listId"`
	Completed bool   `json:"completed"`
}

type CreateTask struct {
	Text      string `json:"text" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}

type UpdateTaskInput struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
