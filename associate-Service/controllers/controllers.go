package controllers

import (
	"github.com/VJ-Vijay77/associateServiceMiniProject/initializers"
	"github.com/VJ-Vijay77/associateServiceMiniProject/middleware"
	"github.com/VJ-Vijay77/associateServiceMiniProject/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Data struct {
	Db *gorm.DB
}

func InitCntlrs() *Data {
	return &Data{
		Db: initializers.DB,
	}
}

func (d *Data) Login(c *gin.Context) {
	var credentials []models.Associate
	var user models.Associate

	d.Db.Find(&credentials)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, "Couldnt bind json")
		return
	}

	for _, i := range credentials {
		if i.Phone == user.Phone && i.Password == user.Password {
			token := middleware.CreateSession(int(i.ID), c)
			c.JSON(200, gin.H{
				"message":       "Login Successfull",
				"session token": token,
			})
			return
		}
	}
	c.JSON(401, gin.H{
		"status": "Failed to login",
		"hint":   "Try to change your username or password",
	})
}

func (d *Data) Dashboard(c *gin.Context) {
	var complaints []models.Complaint
	t := d.Db.Find(&complaints)
	if t.Error != nil {
		c.JSON(400, "Couldnt fetch from database")
		return
	}

	c.JSON(200, complaints)
}

func (d *Data) Home(c *gin.Context) {
	if !middleware.IsExpired(c) {
		c.JSON(200, "Welcome to Home, Please login")
		return
	}

	c.JSON(200, gin.H{
		"message": "Welcome to Home",
	})
}

func (d *Data) Status(c *gin.Context) {
	token := c.Param("token")
	var status models.Complaint
	t := d.Db.Select("tokenno", "brand", "model", "complaints", "status", "associateid","complaintdate","resolvedate").Where("tokenno=?", token).Find(&status)
	if t.Error != nil || status.Tokenno == 0 {
		c.JSON(400, "Couldnt get status , invalid token number")
		return
	}
	c.JSON(200, gin.H{
		"tokenno":     status.Tokenno,
		"brand":     status.Brand,
		"model":     status.Model,
		"complaint": status.Complaints,
		"status":    status.Status,
		"associateid": status.Associateid,
		"complaintdate":status.Complaintdate,
		"resolvedate":status.Resolvedate,
	})

}
