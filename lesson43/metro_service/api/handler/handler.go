package handler

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/storage/postgres"
)

type handler struct {
	Station *postgres.StationRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		Station: postgres.NewStationRepo(db),
	}
}
