package dbcon

import "database/sql"

type EnrollmentRepo struct {
	Db *sql.DB
}

func ConnectEnrollment(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{Db: db}
}
