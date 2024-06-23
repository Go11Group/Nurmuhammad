package postgres

import (
	"database/sql"
	"fmt"
	"user/models"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (s *UserRepo) Create(user *models.User) error {

	_, err := s.Db.Exec("insert into users(name,phone,age) values ($1,$2,$3)",
		user.Name, user.Phone, user.Age)
	return err
}

func (s *UserRepo) GetById(id string) (*models.User, error) {
	var user = models.User{Id: id}

	err := s.Db.QueryRow("select id,name,phone,age from users where id = $1", id).
		Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserRepo) Delete(id string) error {
	_, err := s.Db.Exec(`DELETE FROM transaction WHERE card_id IN (
		SELECT id FROM card WHERE user_id = $1)`, id)
	if err != nil {
		return err
	}
	_, err = s.Db.Exec("delete from card where user_id = $1", id)
	if err != nil {
		return err
	}

	_, err = s.Db.Exec("delete from users where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	rows, err := s.Db.Query("select id,name,phone,age from users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
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
	if len(user.Phone) > 0 {
		query += fmt.Sprintf("phone=$%d, ", n)
		arr = append(arr, user.Phone)
		n++
	}
	if user.Age > 0 {
		query += fmt.Sprintf("age=$%d, ", n)
		arr = append(arr, user.Age)
		n++
	}
	arr = append(arr, id)

	query = query[:len(query)-2] + fmt.Sprintf(" where user_id=$%d", n)

	_, err := us.Db.Exec(query, arr...)
	if err != nil {
		return nil, err
	}
	newuser := models.User{}
	err = us.Db.QueryRow(`select id,
    name,
    phone
    age from users WHERE id = $1`, id).Scan(
		&newuser.Id, &newuser.Name, &newuser.Phone, &newuser.Age)
	if err != nil {
		return nil, err
	}

	return &newuser, nil
}
