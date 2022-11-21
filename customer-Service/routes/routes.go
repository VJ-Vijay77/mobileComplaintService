package routes

import (
	"github.com/VJ-Vijay77/customerServiceMiniProject/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	c := controllers.InitCntlrs()
	r.POST("/adduser", c.AddUser)
	// r.DELETE("/deleteuser",c.DeleteUser)
	r.POST("/login", c.Login)
	r.POST("/complaints", c.Complaints)
	r.GET("/dashboard", c.Dashboard)
	r.GET("/status/:token", c.ComplaintStatus)
	r.GET("/home", c.Home)
	r.GET("/logout", c.Logout)
}
