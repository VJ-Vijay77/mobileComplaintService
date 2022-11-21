package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/VJ-Vijay77/authServiceMiniProject/controllers"
)
func Routes(r *gin.Engine) {
	c := controllers.InitCntlrs()
	r.POST("/authenticate",c.Authenticate)
	r.GET("/getuser",c.GetAllUsers)
	r.GET("/ping",c.Ping)
}
