package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

type ProblemRepo struct {
	DB *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{DB: db}
}

func (b *ProblemRepo) ProblemsGetById(id string) (model.Problem, error) {
	var prblem model.Problem
	err := b.DB.QueryRow("SELECT * FROM problems WHERE problem_id = $1", id).Scan(
		&prblem.ProblemID,
		&prblem.ProblemName,
		&prblem.Description,
	)
	return prblem, err
}

func (b *ProblemRepo) ProblemDeleteById(id string) error {

	_, err := b.DB.Exec(`DELETE FROM solved_problems WHERE problem_id = $1`, id)
	if err != nil {
		return err
	}
	_, err = b.DB.Exec(`DELETE FROM problems WHERE problem_id = $1`, id)
	return err
}

func (b *ProblemRepo) ProblemUpdateById(prblem *model.Problem, id int) error {

	_, err := b.DB.Exec(`
	UPDATE problems
	SET Problem_name=$1,description=$2
	WHERE problem_id = $3`,
		prblem.ProblemName,
		prblem.Description,
		id)

	return err
}
