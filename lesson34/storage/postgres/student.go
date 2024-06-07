package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{db: db}
}

func (b *StudentRepo) GetById(id string) (model.Student, error) {
	var student model.Student
	err := b.db.QueryRow("SELECT * FROM student WHERE Student_ID = $1", id).Scan(
		&student.StudentID,
		&student.FirstName,
		&student.LastName,
		&student.Email,
		&student.DateOfBirth,
		&student.EnrollmentDate,
	)

	return student, err
}

func (b *StudentRepo) DeleteById(id string) error {

	_, err := b.db.Exec(`DELETE FROM student WHERE student_id = $1`, id)
	return err
}

func (b *StudentRepo) UpdateById(student *model.Student, id int) error {

	_, err := b.db.Exec(`
	UPDATE student
	SET first_name = $1,
		last_name = $2,
		email = $3,
		date_of_birth = $4,
		enrollment_date = $5
	WHERE student_id = $6`,
		student.FirstName,
		student.LastName,
		student.Email,
		student.DateOfBirth,
		student.EnrollmentDate,
		id)

	return err
}

// func (b *StudentRepo) GetAll()
