package dbcon

import "database/sql"

type LessonRepo struct {
	Db *sql.DB
}

func ConnectLesson(db *sql.DB) *LessonRepo {
	return &LessonRepo{Db: db}
}
