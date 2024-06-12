package dbcon

import (
	"database/sql"
	"exam/models"
)

type UserRepo struct {
	Db *sql.DB
}

func ConnectUser(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) GetAll() ([]models.User, error) {

}
