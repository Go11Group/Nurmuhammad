package postgres

import (
	"database/sql"
	// "fmt"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

type SolvedRepo struct {
	db *sql.DB
}

func NewSolvedRepo(db *sql.DB) *SolvedRepo {
	return &SolvedRepo{db: db}
}

func (b *SolvedRepo) GetAll() ([]model.SolvedProblems, error) {
	users := []model.SolvedProblems{}
	rows, err := b.db.Query(`select u.first_name,u.last_name,u.email,p.Problem_name,p.description
	from users as u
	join Solved_problems as s
	on s.user_id=u.user_id
	join Problems as p
	on p.Problem_id=s.Problem_id`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := model.SolvedProblems{}
		err = rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.ProblemName, &user.Description)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}
