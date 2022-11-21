package routes

import (
	"github.com/VJ-Vijay77/associateServiceMiniProject/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	c := controllers.InitCntlrs()
	r.POST("/login", c.Login)
	r.GET("/home",c.Home)
	r.GET("/dashboard",c.Dashboard)
	r.GET("/status/:token",c.Status)
	r.PUT("/update",c.UdpateStatus)
	
}
