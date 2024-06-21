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

type UserRepo struct {
	Db *sql.DB
}

func ConnectUser(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

type Filter struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age1   int    `json:"ageone"`
	Age2   int    `json:"agetwo"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type SearchFilter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age1  int    `json:"ageone"`
	Age2  int    `json:"agetwo"`
}

func (user *UserRepo) GetAll(f Filter) (*models.UserGetAllResp, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)
	countquery := `select count(*) from users `

	query := `select user_id,
    name,
    email,
    birthday,
    password,
    created_at,
    updated_at
	from users `

	filter := ``

	if len(f.UserID) > 0 {
		params["user_id"] = f.UserID
		filter += ` and user_id = :user_id `

	}

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += ` and name = :name `
	}

	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += ` and email = :email `
	}

	if f.Age1 > 0 && f.Age2 > 0 {
		params["age1"] = f.Age1
		params["age2"] = f.Age2
		filter += ` AND EXTRACT(YEAR FROM age(birthday)) between :age1 and :age2 `
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
	err := user.Db.QueryRow(countquery, arr1...).Scan(&count)
	if err != nil {
		return nil, err
	}

	query, arr = replaceQueryParams(query, params)
	rows, err := user.Db.Query(query, arr...)

	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {

		user := models.User{}

		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	count = int(math.Ceil(float64(count) / float64(f.Limit)))
	return &models.UserGetAllResp{Users: users, Count: count}, err

}

func replaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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

func (us *UserRepo) GetUser(id string) (models.User, error) {
	var user models.User
	err := us.Db.QueryRow(`select user_id,
    name,
    email,
    birthday,
    password,
    created_at,
    updated_at from users WHERE deleted_at=0 and user_id = $1`, id).Scan(
		&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (us *UserRepo) InsertToUser(user models.User) error {

	_, err := us.Db.Exec(`insert into users(name,email,birthday,password)
	values ($1,$2,$3,$4)`, &user.Name, &user.Email, &user.Birthday, &user.Password)
	return err
}

func (us *UserRepo) DeleteUser(id string) error {
	var time int
	err := us.Db.QueryRow(`select deleted_at from users where user_id = $1`, id).Scan(&time)
	if err != nil {
		return err
	}
	if time != 0 {
		return errors.New("this user is already deleted")
	}

	_, err = us.Db.Exec(`update users set
	deleted_at = date_part('epoch', current_timestamp)::INT
   where user_id = $1 and deleted_at = 0`, id)
	return err
}

func (us *UserRepo) UpdateUser(id string, user models.User) (*models.User, error) {
	query := `update users set `
	n := 1
	var arr []interface{}
	if len(user.Name) > 0 {
		query += fmt.Sprintf("name=$%d, ", n)
		arr = append(arr, user.Name)
		n++
	}
	if len(user.Email) > 0 {
		query += fmt.Sprintf("email=$%d, ", n)
		arr = append(arr, user.Email)
		n++
	}
	if len(user.Birthday) > 0 {
		query += fmt.Sprintf("birthday=$%d, ", n)
		arr = append(arr, user.Birthday)
		n++
	}
	if len(user.Password) > 0 {
		query += fmt.Sprintf("password=$%d, ", n)
		arr = append(arr, user.Password)
		n++
	}
	arr = append(arr, id)

	query += fmt.Sprintf("updated_at=current_timestamp where user_id=$%d and deleted_at=0", n)

	_, err := us.Db.Exec(query, arr...)
	if err != nil {
		return nil, err
	}
	newuser := models.User{}
	err = us.Db.QueryRow(`select user_id,
    name,
    email,
    birthday,
    password,
    created_at,
    updated_at from users WHERE deleted_at=0 and user_id = $1`, id).Scan(
		&newuser.UserID, &newuser.Name, &newuser.Email, &newuser.Birthday, &newuser.Password, &newuser.CreatedAt, &newuser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &newuser, nil
}

func (us *UserRepo) GetUserCourses(id string) (*models.UserCourses, error) {
	users := models.UserCourses{UserId: id}
	rows, err := us.Db.Query(`
	SELECT c.course_id, c.title, c.description
	FROM courses c
	JOIN enrollments e ON c.course_id = e.course_id
	WHERE e.user_id = $1 AND c.deleted_at = 0
`, id)

	for rows.Next() {

		course := models.Course2{}
		err = rows.Scan(&course.CourseID, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		users.Courses = append(users.Courses, course)
	}

	return &users, err
}

func (us *UserRepo) SearchUsers(f SearchFilter) (*models.SearchUser, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)

	query := `select user_id,
    name,
    email
	from users `

	filter := ``

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += ` and name = :name `
	}

	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += ` and email = :email `
	}

	if f.Age1 > 0 && f.Age2 > 0 {
		params["age1"] = f.Age1
		params["age2"] = f.Age2
		filter += ` AND EXTRACT(YEAR FROM age(birthday)) between :age1 and :age2 `
	}

	if len(filter) > 0 {
		query = query + ` where deleted_at=0 ` + filter
	}

	query, arr = replaceQueryParams(query, params)
	rows, err := us.Db.Query(query, arr...)

	if err != nil {
		return nil, err
	}

	users := models.SearchUser{}
	for rows.Next() {

		user := models.User2{}

		err = rows.Scan(&user.UserID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users.Results = append(users.Results, user)
	}

	return &users, err
}
