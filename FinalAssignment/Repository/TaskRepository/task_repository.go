package TaskRepository

import (
	models "FinalAssignment/Repository/Models"
	"gorm.io/gorm"
)

func GetAllTasks(db *gorm.DB) ([]models.Tasks, error) {
	var task []models.Tasks
	if err := db.Find(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func FindTaskById(id int, db *gorm.DB) (models.Tasks, bool, error) {
	var taskToReturn models.Tasks

	if err := db.Where("id = ?", id).First(&taskToReturn).Error; err != nil {
		return taskToReturn, false, err
	}

	return taskToReturn, true, nil
}

func UpdateTask(id int, db *gorm.DB, task *models.Tasks) error {

	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		return err
	}

	if err := db.Save(&task).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTask(id int, db *gorm.DB) error {
	var task models.Tasks
	if err := db.Where("id = ?", id).Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAllTasks(db *gorm.DB) {
	db.Exec(`DELETE FROM tasks`)
}
