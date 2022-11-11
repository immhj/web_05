package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_05/dao"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	}
	dao.AddUSer(username, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register success",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user does not exists",
		})
		return
	}
	selectpassword := dao.SelectPasswordFromUsername(username)
	if selectpassword != password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
	})
}

func Updatapassword(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	dao.Updatepassword(username, password)
	if password == dao.SelectPasswordFromUsername(username) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail to updata password",
		})
	}
}
