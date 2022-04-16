package main

import (
	"gorm.io/gorm"
)

type StoriesRepo struct {
	db *gorm.DB
}

func NewStoriesRepo(db *gorm.DB) *StoriesRepo {
	return &StoriesRepo{db: db}
}

func (timeFromDB *StoriesRepo) GetTime() string {
	var result topstories
	if timeFromDB.db.First(&result) != nil {
		return result.TimeStamp
	}

	return ""
}

func (timeFromDB *StoriesRepo) DeleteAll() {
	timeFromDB.db.Exec(`DELETE FROM topstories`)
}

func (insertData *StoriesRepo) InsertData(insert []*topstories) {
	insertData.db.Create(&insert)
}

func (getAll *StoriesRepo) ReadAll(stories []*topstories) {

	getAll.db.Find(&stories)
}
