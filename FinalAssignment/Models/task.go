package Models

type Task struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Text      string `json:"text"`
	ListId    int    `json:"listId"`
	Completed bool   `json:"completed"`
}

type CreateTask struct {
	Text string `json:"text" binding:"required"`
	//ListId int    `json:"listId" binding:"required"`
}

type UpdateTaskInput struct {
	Text string `json:"text"`
}
