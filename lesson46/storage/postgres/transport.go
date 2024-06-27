package postgres

import (
	"database/sql"
	"new/genproto/transportService"
)

type TransportRepo struct {
	Db *sql.DB
}

func ConnectBook(db *sql.DB) *TransportRepo {
	return &TransportRepo{Db: db}
}

func (t *TransportRepo) GetBusSchedule(num *transportService.BusNumber) (*transportService.GetBusScheduleResponse, error) {
	bus := transportService.GetBusScheduleResponse{Busnum: num.Busnum}
	err := t.Db.QueryRow(`select schedule from transport where number=$1`, num.Busnum).Scan(&bus.Places)
	return &bus, err
}

func (t *TransportRepo) GetBusLocation(num *transportService.BusNumber) (*transportService.TrackBusLocationResponse, error) {
	bus := transportService.TrackBusLocationResponse{Busnum: num.Busnum}
	err := t.Db.QueryRow(`select location from transport where number=$1`, num.Busnum).Scan(&bus.Location)
	return &bus, err
}

func (t *TransportRepo) UpdateFeedback(num *transportService.ReportTrafficJamRequest) error {
	_, err := t.Db.Exec(`update transport set feedback=$1 where number=$2`, num.Feedback, num.Busnum)
	return err
}
