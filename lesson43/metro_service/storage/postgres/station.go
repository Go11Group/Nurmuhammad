package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson43/metro_service/models"
	"github.com/google/uuid"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) Create(station *models.CreateStation) error {

	_, err := s.Db.Exec("insert into station(id, name) values ($1, $2)",
		uuid.NewString(), station.Name)

	return err
}

func (s *StationRepo) GetById(id string) (*models.Station, error) {
	var station = models.Station{Id: id}

	err := s.Db.QueryRow("select name from station where id = $1", id).
		Scan(&station.Name)
	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (s *StationRepo) Delete(id string) error {
	_, err := s.Db.Exec("delete from terminal where station_id = $1", id)
	if err != nil {
		return err
	}
	_, err = s.Db.Exec("delete from station where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *StationRepo) Update(id string, m *models.CreateStation) error {
	_, err := s.Db.Exec("Update station set name=$1 where id = $2", m.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *StationRepo) GetAll() ([]models.Station, error) {
	var stations []models.Station
	rows, err := s.Db.Query("select id,name from station")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		station := models.Station{}
		err = rows.Scan(&station.Id, &station.Name)
		if err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	return stations, nil
}
