package router

import (
	"gin_curd/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", api.Users)
	r.POST("/user", api.CreateUser)
	r.PUT("/user/:id", api.UpdateUser)
	r.DELETE("/user/:id", api.DeleteUser)

	return r
}
