package main

import (
	"new/api"
	"new/api/handler"
	dbcon "new/storage/postgres"

	"github.com/casbin/casbin/v2"
)

func main() {
	en, err := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	if err != nil {
		panic(err)
	}
	db, err := dbcon.ConnnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	user := dbcon.NewUserRepository(db)
	hand := NewHandler(user, en)
	c := api.ConnectGin(hand)
	c.Run()
}
func NewHandler(user *dbcon.UserRepo, En *casbin.Enforcer) *handler.Handler {
	return &handler.Handler{
		User:     user,
		Enforcer: En,
	}
}
