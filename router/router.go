package router

import (
	"gin_curd/router/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	router.LoadUser(r)
	return r
}
