package controllers

import (
	"encoding/json"
	"io"

	"github.com/VJ-Vijay77/authServiceMiniProject/initializers"
	"github.com/VJ-Vijay77/authServiceMiniProject/models"
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

func (d *Data) Authenticate(c *gin.Context) {
	var data models.User

	da, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(da, &data)
	
	var user models.User
	
	d.Db.Where("phone=?",data.Phone).Find(&user)
	
	if user.Phone == data.Phone && user.Password == data.Password {
		c.JSON(200, user.ID)
		return
	}
	
	c.JSON(401,gin.H{
		"status": "wrong credentials",
	})
}

func (d *Data) Ping(c *gin.Context) {
	c.JSON(200, "The connectino is fine...")
}

func (d *Data) GetAllUsers(c *gin.Context) {
	var data []models.User
	d.Db.Find(&data)
	c.JSON(200, data)
}
