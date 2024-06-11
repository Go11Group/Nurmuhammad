package handler

import (
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User    *postgres.UserRepo
	Problem *postgres.ProblemRepo
	Solved  *postgres.SolvedRepo
}

func StartGin(handler Handler) (c *gin.Engine) {
	m := gin.Default()
	m.Any("/user/:id", handler.user)
	m.Any("/problem/:id", handler.problem)
	m.GET("/solvedproblem", handler.GetAllSolved)
	m.Run()
	return m
}
