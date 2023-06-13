package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"record_code/app/service"
)

type UserControllerInterface interface {
	GetByID(c *gin.Context)
}

type UserController struct {
}

func (u *UserController) GetByID(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data":    gin.H{},
		})
		return
	}

	users := service.GetDetailById(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data":    users,
		"code":    http.StatusOK,
	})
}
