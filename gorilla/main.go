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
	stu := postgres.NewStudentRepo(db)

	server := handler.NewHandler(handler.Handler{Student: stu})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
