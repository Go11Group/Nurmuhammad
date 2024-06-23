package main

import (
	"github.com/Go11Group/at_lesson/lesson43/metro_service/api"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := api.Routes(db)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
