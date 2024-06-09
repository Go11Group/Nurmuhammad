package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson34/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (b *UserRepo) GetById(id string) (model.User, error) {
	var user model.User
	err := b.db.QueryRow("SELECT * FROM users WHERE user_ID = $1", id).Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EnrollmentDate,
	)
	return user, err
}

func (b *UserRepo) DeleteById(id string) error {

	_, err := b.db.Exec(`DELETE FROM users WHERE user_id = $1`, id)
	return err
}

func (b *UserRepo) UpdateById(user *model.User, id int) error {

	_, err := b.db.Exec(`
	UPDATE users
	SET first_name = $1,
		last_name = $2,
		email = $3,
		enrollment_date = $5
	WHERE user_id = $6`,
		user.FirstName,
		user.LastName,
		user.Email,
		user.EnrollmentDate,
		id)

	return err
}
