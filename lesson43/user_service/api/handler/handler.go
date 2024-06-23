package handler

import (
	"database/sql"
	"user/storage/postgres"
)

type handler struct {
	User *postgres.UserRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		User: postgres.NewUserRepo(db),
	}
}
