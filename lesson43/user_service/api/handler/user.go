package handler

import (
	"fmt"
	"net/http"
	"user/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateUser(ctx *gin.Context) {
	stn := models.User{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.User.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param(":id")
	err := h.User.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param(":id")
	user := models.User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users, err := h.User.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, users)

}

func (h *handler) GetByIdUser(ctx *gin.Context) {
	id := ctx.Param(":id")

	user, err := h.User.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (h *handler) GetAllUser(ctx *gin.Context) {
	users, err := h.User.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, users)
}
