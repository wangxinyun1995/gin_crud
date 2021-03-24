package router

import (
	"gin_curd/api"
	"github.com/gin-gonic/gin"
)

func LoadUser(e *gin.Engine) {
	e.GET("/users", api.Users)
	e.POST("/user", api.CreateUser)
	e.PUT("/user/:id", api.UpdateUser)
	e.DELETE("/user/:id", api.DeleteUser)
}