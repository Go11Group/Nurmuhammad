package handler

import (
	"github.com/casbin/casbin/v2"
	dbcon "new/storage/postgres"
)

type Handler struct {
	User     *dbcon.UserRepo
	Enforcer *casbin.Enforcer
}
