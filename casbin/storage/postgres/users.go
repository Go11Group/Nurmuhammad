package dbcon

import (
	"database/sql"
	"fmt"
	"new/structs"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) GetUserById(id string) (*structs.User, error) {
	query := `SELECT id, username, email FROM users WHERE id = $1`

	var user structs.User

	err := u.Db.QueryRow(query, id).Scan(&user.UserID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("not user in this id %s", id)
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) DeleteUserById(id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := u.Db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (u *UserRepo) GetUser(user *structs.UserInfo) (*structs.UserInfoRes, error) {
	query := `SELECT id, role FROM users WHERE email = $1 AND password = $2`

	var userInfoRes structs.UserInfoRes

	err := u.Db.QueryRow(query, user.Email, user.Password).Scan(&userInfoRes.UserID, &userInfoRes.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email %s and the provided password", user.Email)
		}
		return nil, err
	}
	return &userInfoRes, nil
}
