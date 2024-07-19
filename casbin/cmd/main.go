package main

import (
	"new/api"
	"new/api/handler"
	dbcon "new/storage/postgres"
)

func main() {
	db, err := dbcon.ConnnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	user := dbcon.NewUserRepository(db)
	hand := NewHandler(user)
	c := api.ConnectGin(hand)
	c.Run()
}
func NewHandler(user *dbcon.UserRepo) *handler.Handler {
	return &handler.Handler{
		User: user,
	}
}
