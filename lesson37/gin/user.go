package gin

import (
	"net/http"

	"github.com/Go11Group/at_lesson/lesson37/psql"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUser(c *gin.Context) {

	filter := psql.Filter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := h.User.GetAll(filter)
	if err != nil {
		c.Writer.Write([]byte("error on read users"))
	}
	for _, v := range users {
		c.JSON(http.StatusBadRequest, v)
	}

}
