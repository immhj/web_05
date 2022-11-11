package apis

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", Login)
	r.POST("/update", Updatapassword)
	r.Run()
}
