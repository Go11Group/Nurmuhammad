package main

import (
	"github.com/Go11Group/at_lesson/lesson34/handler"
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	stu := postgres.NewUserRepo(db)
	pr := postgres.NewProblemRepo(db)
	sol := postgres.NewSolvedRepo(db)
	all := handler.Handler{Solved: sol, User: stu, Problem: pr}
	handler.StartGin(all)
}
