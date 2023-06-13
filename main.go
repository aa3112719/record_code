package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"record_code/app/controller"
	"record_code/utils/psql"
)

func main() {
	r := gin.Default()
	userController := new(controller.UserController)

	defer psql.DB.Close()
	r.GET("/ping", userController.GetByID)
	r.Run()
}
