package ListRepository

import (
	models "FinalAssignment/Repository/Models"
	"gorm.io/gorm"
)

func FindAllLists(db *gorm.DB) ([]models.Lists, error) {
	var list []models.Lists
	if err := db.Find(&list).Error; err != nil {
		return list, err
	}

	return list, nil
}

func GetListById(id int, db *gorm.DB) (models.Lists, bool, error) {
	list := models.Lists{}

	if err := db.Where("id = ?", id).First(&list).Error; err != nil {
		return list, false, err
	}

	return list, true, nil
}

func DeleteList(id int, db *gorm.DB) error {
	var list models.Lists
	var task models.Tasks

	if err := db.Where("id = ?", id).Delete(&list).Error; err != nil {
		return err
	}

	if err := db.Where("list_id = ?", id).Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
