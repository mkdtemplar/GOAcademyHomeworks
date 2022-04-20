package Lecture_27_Task

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type StoriesRepo struct {
	*sql.DB
}

func NewStoriesRepo(DB *sql.DB) *StoriesRepo {
	return &StoriesRepo{DB: DB}
}
