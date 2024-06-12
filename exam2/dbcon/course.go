package dbcon

import "database/sql"

type CourseRepo struct {
	Db *sql.DB
}

func ConnectCourse(db *sql.DB) *CourseRepo {
	return &CourseRepo{Db: db}
}
