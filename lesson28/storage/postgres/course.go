package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/google/uuid"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course *model.Course) error {
	_, err := c.DB.Exec(`insert into course(id,name,field) values ($1,$2,$3)`,
		uuid.NewString(), course.Name, course.Field)
	if err != nil {
		return err
	}

	return nil
}

func (c *CourseRepo) Update(course *model.Course) error {

	_, err := c.DB.Exec(`Update course
	set name=$1,field=$2
	where id=$3`,
		course.Name, course.Field, course.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c *CourseRepo) GetAllStudents() ([]model.Course, error) {
	rows, err := c.DB.Query(`select id,name,field from course`)
	if err != nil {
		return nil, err
	}

	var courses []model.Course
	var course model.Course
	for rows.Next() {
		err = rows.Scan(&course.Id, &course.Name, &course.Field)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (c *CourseRepo) GetByID(id string) (model.Course, error) {
	var course model.Course

	err := c.DB.QueryRow(`select id,name,field from course where id=$1`, id).
		Scan(&course.Id, &course.Name, &course.Field)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}
