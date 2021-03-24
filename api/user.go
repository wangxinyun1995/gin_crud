package api

import (
	"gin_curd/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	var user models.User
	result, err := user.GetUsers()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "not find",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.PostForm("password")
	re, err := user.Update(id)
	if err != nil || re.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": re,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": re,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	re, err := user.DestroyUser(id)
	if err != nil || re.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": re,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": re,
	})

}
