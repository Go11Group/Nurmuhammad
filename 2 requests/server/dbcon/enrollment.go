package dbcon

import (
	"database/sql"
	"errors"
	"exam/models"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type EnrollmentRepo struct {
	Db *sql.DB
}

func ConnectEnrollment(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{Db: db}
}

type FilterEnrollment struct {
	EnrollmentID string `json:"enrollmentId"`
	UserID       string `json:"userId"`
	CourseID     string `json:"courseId"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

func (en *EnrollmentRepo) GetAllEnrollment(f FilterEnrollment) (*models.EnrollmentGetAllResp, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)
	countquery := `select count(*) from enrollments `

	query := `select 
	enrollment_id,
	user_id,
	course_id,
	enrollment_date,
	created_at,
	updated_at
	from enrollments `

	filter := ``

	if len(f.EnrollmentID) > 0 {
		params["enrollment_id"] = f.EnrollmentID
		filter += ` and enrollment_id = :enrollment_id `

	}

	if len(f.UserID) > 0 {
		params["user_id"] = f.UserID
		filter += ` and user_id = :user_id `
	}

	if len(f.CourseID) > 0 {
		params["course_id"] = f.CourseID
		filter += ` and course_id = :course_id `
	}

	limit = ""
	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if f.Offset > 0 {
		params["offset"] = f.Offset
		limit += ` OFFSET :offset`
	}

	if len(filter) > 0 || len(limit) > 0 {
		query = query + ` where deleted_at=0 ` + filter + limit
		countquery = countquery + ` where deleted_at=0 ` + filter
	}

	count := 0
	countquery, arr1 := replaceQueryParams(countquery, params)
	err := en.Db.QueryRow(countquery, arr1...).Scan(&count)
	if err != nil {
		return nil, err
	}

	query, arr = replaceQueryParamsEnrollment(query, params)
	rows, err := en.Db.Query(query, arr...)

	if err != nil {
		return nil, err
	}

	enrollments := []models.Enrollment{}
	for rows.Next() {

		enrollment := models.Enrollment{}

		err = rows.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	count = int(math.Ceil(float64(count) / float64(f.Limit)))
	return &models.EnrollmentGetAllResp{Enrollment: enrollments, Count: count}, err

}

func replaceQueryParamsEnrollment(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {

		if k != "" && strings.Contains(namedQuery, ":"+k) {

			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}

func (en *EnrollmentRepo) GetEnrollment(id string) (models.Enrollment, error) {
	var enrollment models.Enrollment
	err := en.Db.QueryRow(`select 
	enrollment_id,
	user_id,
	course_id,
	enrollment_date,
	created_at,
	updated_at from enrollments WHERE deleted_at=0 and enrollment_id = $1`, id).Scan(
		&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt)
	return enrollment, err
}

func (en *EnrollmentRepo) InsertToEnrollment(enrollment models.Enrollment) error {

	_, err := en.Db.Exec(`insert into enrollments(user_id,course_id,enrollment_date)
	values ($1,$2,$3)`, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate)
	return err
}

func (en *EnrollmentRepo) DeleteEnrollment(id string) error {
	var time int
	err := en.Db.QueryRow(`select deleted_at from enrollments where enrollment_id = $1`, id).Scan(&time)
	if err != nil {
		return err
	}
	if time != 0 {
		return errors.New("this enrollment is already deleted")
	}

	_, err = en.Db.Exec(`update enrollments set
	deleted_at = date_part('epoch', current_timestamp)::INT
   where enrollment_id = $1 and deleted_at = 0`, id)
	return err
}

func (en *EnrollmentRepo) UpdateEnrollment(id string, enrollment models.Enrollment) (*models.Enrollment, error) {
	query := `update enrollments set `
	n := 1
	var arr []interface{}
	if len(enrollment.EnrollmentID) > 0 {
		query += fmt.Sprintf("enrollment_id=$%d, ", n)
		arr = append(arr, enrollment.EnrollmentID)
		n++
	}
	if len(enrollment.UserID) > 0 {
		query += fmt.Sprintf("user_id=$%d, ", n)
		arr = append(arr, enrollment.UserID)
		n++
	}
	if len(enrollment.CourseID) > 0 {
		query += fmt.Sprintf("course_id=$%d, ", n)
		arr = append(arr, enrollment.CourseID)
		n++
	}

	arr = append(arr, id)

	query += fmt.Sprintf("updated_at=current_timestamp where enrollment_id=$%d and deleted_at=0", n)

	_, err := en.Db.Exec(query, arr...)
	if err != nil {
		return nil, err
	}
	newenrollment := models.Enrollment{}
	err = en.Db.QueryRow(`select enrollment_id,
	user_id,
	course_id,
	enrollment_date,
	created_at,
	updated_at
	from enrollments WHERE deleted_at=0 and enrollment_id = $1`, id).Scan(&newenrollment.EnrollmentID, &newenrollment.UserID, &newenrollment.CourseID, &newenrollment.EnrollmentDate, &newenrollment.CreatedAt, &newenrollment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &newenrollment, nil
}
