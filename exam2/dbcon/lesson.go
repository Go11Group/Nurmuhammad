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

type LessonRepo struct {
	Db *sql.DB
}

func ConnectLesson(db *sql.DB) *LessonRepo {
	return &LessonRepo{Db: db}
}

type FilterLesson struct {
	LessonID string `json:"lessonId"`
	CourseID string `json:"courseId"`
	Title    string `json:"title"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

func (ls *LessonRepo) GetAllLesson(f FilterLesson) (*models.LessonGetAllResp, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)
	countquery := `select count(*) from lessons `

	query := `select lesson_id,
    course_id,
    title,
    content,
    created_at,
    updated_at
	from lessons `

	filter := ``

	if len(f.LessonID) > 0 {
		params["lesson_id"] = f.LessonID
		filter += ` and lesson_id = :lesson_id `

	}

	if len(f.CourseID) > 0 {
		params["course_id"] = f.CourseID
		filter += ` and course_id = :course_id `
	}

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += ` and title = :title `
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

	err := ls.Db.QueryRow(countquery, arr1...).Scan(&count)
	if err != nil {
		return nil, err
	}

	query, arr = replaceQueryParamslesson(query, params)
	rows, err := ls.Db.Query(query, arr...)

	if err != nil {
		return nil, err
	}

	lessons := []models.Lesson{}
	for rows.Next() {

		lesson := models.Lesson{}

		err = rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	count = int(math.Ceil(float64(count) / float64(f.Limit)))
	return &models.LessonGetAllResp{Lesson: lessons, Count: count}, err

}

func replaceQueryParamslesson(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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

func (ls *LessonRepo) GetLesson(id string) (models.Lesson, error) {
	var lesson models.Lesson
	err := ls.Db.QueryRow(`select lesson_id,
    course_id,
    title,
    content,
    created_at,
    updated_at from lessons WHERE deleted_at=0 and lesson_id = $1`, id).Scan(
		&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
	return lesson, err
}

func (ls *LessonRepo) InsertToLesson(lesson models.Lesson) error {

	_, err := ls.Db.Exec(`insert into lessons(course_id,title,content)
	values ($1,$2,$3)`, &lesson.CourseID, &lesson.Title, &lesson.Content)
	return err
}

func (ls *LessonRepo) DeleteLesson(id string) error {
	var time int
	err := ls.Db.QueryRow(`select deleted_at from lessons where lesson_id = $1`, id).Scan(&time)
	if err != nil {
		return err
	}
	if time != 0 {
		return errors.New("this lesson is already deleted")
	}

	_, err = ls.Db.Exec(`update lessons set
	deleted_at = date_part('epoch', current_timestamp)::INT
   where lesson_id = $1 and deleted_at = 0`, id)
	return err
}

func (ls *LessonRepo) UpdateLesson(id string, lesson models.Lesson) (*models.Lesson, error) {
	query := `update lessons set `
	n := 1
	var arr []interface{}
	if len(lesson.CourseID) > 0 {
		query += fmt.Sprintf("course_id=$%d, ", n)
		arr = append(arr, lesson.CourseID)
		n++
	}
	if len(lesson.Title) > 0 {
		query += fmt.Sprintf("title=$%d, ", n)
		arr = append(arr, lesson.Title)
		n++
	}
	if len(lesson.Content) > 0 {
		query += fmt.Sprintf("content=$%d, ", n)
		arr = append(arr, lesson.Content)
		n++
	}

	arr = append(arr, id)

	query += fmt.Sprintf("updated_at=current_timestamp where lesson_id=$%d and deleted_at=0", n)

	_, err := ls.Db.Exec(query, arr...)
	if err != nil {
		return nil, err
	}
	newlesson := models.Lesson{}
	err = ls.Db.QueryRow(`select lesson_id,
	course_id,
    title,
    content,
    created_at,
    updated_at
	from lessons WHERE deleted_at=0 and lesson_id = $1`, id).Scan(
		&newlesson.LessonID, &newlesson.CourseID, &newlesson.Title, &newlesson.Content, &newlesson.CreatedAt, &newlesson.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &newlesson, nil
}
