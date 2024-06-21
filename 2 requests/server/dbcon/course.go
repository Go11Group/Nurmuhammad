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

type CourseRepo struct {
	Db *sql.DB
}

func ConnectCourse(db *sql.DB) *CourseRepo {
	return &CourseRepo{Db: db}
}

type FilterCourse struct {
	CourseID    string `json:"courseId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

func (cu *CourseRepo) GetAllCourse(f FilterCourse) (*models.CourseGetAllResp, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)
	countquery := `select count(*) from courses `

	query := `select 
	course_id,
	title,
	description,
	created_at,
	updated_at
	from courses `

	filter := ``

	if len(f.CourseID) > 0 {
		params["course_id"] = f.CourseID
		filter += ` and course_id = :course_id `

	}

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += ` and title = :title `
	}

	if len(f.Description) > 0 {
		params["description"] = f.Description
		filter += ` and description = :description `
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
	err := cu.Db.QueryRow(countquery, arr1...).Scan(&count)
	if err != nil {

		return nil, err
	}

	query, arr = replaceQueryParamsCourse(query, params)
	rows, err := cu.Db.Query(query, arr...)

	if err != nil {

		return nil, err
	}

	courses := []models.Course{}
	for rows.Next() {

		course := models.Course{}

		err = rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	count = int(math.Ceil(float64(count) / float64(f.Limit)))
	return &models.CourseGetAllResp{Course: courses, Count: count}, err

}

func replaceQueryParamsCourse(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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

func (cu *CourseRepo) GetCourse(id string) (models.Course, error) {
	var course models.Course
	err := cu.Db.QueryRow(`select 
	course_id,
	title,
	description,
	created_at,
	updated_at from courses WHERE deleted_at=0 and course_id = $1`, id).Scan(
		&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
	return course, err
}

func (cu *CourseRepo) InsertToCourse(course models.Course) error {
	_, err := cu.Db.Exec(`insert into courses(title,description)
	values ($1,$2)`, &course.Title, &course.Description)
	return err
}

func (cu *CourseRepo) DeleteCourse(id string) error {
	var time int
	err := cu.Db.QueryRow(`select deleted_at from courses where course_id = $1`, id).Scan(&time)
	if err != nil {
		return err
	}
	if time != 0 {
		return errors.New("this course is already deleted")
	}

	_, err = cu.Db.Exec(`update courses set
	deleted_at = date_part('epoch', current_timestamp)::INT
   where course_id = $1 and deleted_at = 0`, id)
	return err
}

func (cu *CourseRepo) UpdateCourse(id string, course models.Course) (*models.Course, error) {
	query := `update courses set `
	n := 1
	var arr []interface{}
	if len(course.CourseID) > 0 {
		query += fmt.Sprintf("couse_id=$%d, ", n)
		arr = append(arr, course.CourseID)
		n++
	}
	if len(course.Title) > 0 {
		query += fmt.Sprintf("title=$%d, ", n)
		arr = append(arr, course.Title)
		n++
	}
	if len(course.Description) > 0 {
		query += fmt.Sprintf("course_id=$%d, ", n)
		arr = append(arr, course.Description)
		n++
	}

	arr = append(arr, id)

	query += fmt.Sprintf("updated_at=current_timestamp where course_id=$%d and deleted_at=0", n)

	_, err := cu.Db.Exec(query, arr...)
	if err != nil {
		return nil, err
	}
	newcourse := models.Course{}
	err = cu.Db.QueryRow(`select 
	course_id,
	title,
	description,
	created_at,
	updated_at
	from courses WHERE deleted_at=0 and course_id = $1`, id).Scan(&newcourse.CourseID, &newcourse.Title, &newcourse.Description, &newcourse.CreatedAt, &newcourse.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &newcourse, nil
}

func (cu *CourseRepo) GetCourseLessons(id string) (*models.CourseLessons, error) {
	rows, err := cu.Db.Query(`
        SELECT lesson_id, title, content 
        FROM lessons 
        WHERE course_id = $1 AND deleted_at = 0`, id)
	if err != nil {
		return nil, err
	}
	course := models.CourseLessons{CourseId: id}

	for rows.Next() {
		var lessons models.Lesson2
		err = rows.Scan(&lessons.LessonID, &lessons.Title, &lessons.Content)
		if err != nil {
			return nil, err
		}
		course.Lessons = append(course.Lessons, lessons)
	}
	return &course, nil
}

func (cu *CourseRepo) GetCourseUsers(id string) (*models.CourseUsers, error) {
	rows, err := cu.Db.Query(`SELECT u.user_id, u.name, u.email
	FROM users u
	JOIN enrollments e ON u.user_id = e.user_id
	WHERE e.course_id = $1 AND u.deleted_at = 0 AND e.deleted_at = 0`, id)
	if err != nil {
		return nil, err
	}
	course := models.CourseUsers{CourseId: id}

	for rows.Next() {
		var users models.User2
		err = rows.Scan(&users.UserID, &users.Name, &users.Email)
		if err != nil {
			return nil, err
		}
		course.Users = append(course.Users, users)
	}
	return &course, nil
}

func (cu *CourseRepo) GetPopularCourses(time models.TimePeriod) (*models.ResponseData, error) {
	rows, err := cu.Db.Query(`
	SELECT c.course_id, c.title AS course_title, COUNT(e.enrollment_id) AS enrollments_count
FROM enrollments e
JOIN courses c ON e.course_id = c.course_id
WHERE e.enrollment_date BETWEEN $1 AND $2
GROUP BY c.course_id, c.title
ORDER BY enrollments_count DESC
LIMIT 10 `, time.StartDate, time.EndDate)
	if err != nil {
		return nil, err
	}
	responseData := models.ResponseData{TimePeriod: time}

	for rows.Next() {
		course := models.PopularCourse{}

		err := rows.Scan(&course.CourseID, &course.CourseTitle, &course.EnrollmentsCount)
		if err != nil {
			return nil, err
		}

		responseData.PopularCourses = append(responseData.PopularCourses, course)
	}

	return &responseData, nil

}
