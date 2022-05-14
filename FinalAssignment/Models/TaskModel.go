package Models

type Tasks struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Text      string `json:"text"`
	ListId    int    `json:"list_Id"`
	Completed bool   `json:"completed"`
}

type CreateTask struct {
	Text      string `json:"text" binding:"required"`
	ListId    int    `json:"list_id" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}

type UpdateTaskInput struct {
	//Text string `json:"text"`
	//ListId    int    `json:"list_id"`
	Completed bool `json:"completed"`
}
