package handler

import dbcon "new/storage/postgres"

type Handler struct {
	User *dbcon.UserRepo
}
