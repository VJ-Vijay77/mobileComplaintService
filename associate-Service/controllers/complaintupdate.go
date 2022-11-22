package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/associateServiceMiniProject/middleware"
	"github.com/VJ-Vijay77/associateServiceMiniProject/models"
	"github.com/gin-gonic/gin"
)

func (d *Data) UdpateStatus(c *gin.Context) {
	ok := middleware.IsExpired(c)
	if !ok {
		c.JSON(401, "Please login to continue")
		return
	}
	cook, _ := c.Request.Cookie("session_token")
	token := cook.Value
	userSession := middleware.Sessions[token]
	ID := fmt.Sprintf("%d", userSession.Userid)
	Id, _ := strconv.Atoi(ID)
	var data models.Complaint
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, "Couldnt bing json")
		return
	}
	Resolvedate := time.Now().Format("01-02-2006")

	var complaints models.Complaint
	// d.Db.Model(&models.Complaint{}).Updates(models.Complaint{Status: data.Status,Resolvedate:Resolvedate,Associateid:  uint(Id)})
	t := d.Db.Raw("UPDATE complaints SET status=?, resolvedate=?, associateid=? WHERE tokenno=?",data.Status,Resolvedate,Id,data.Tokenno).Scan(&complaints)
	if t.Error != nil {
		c.JSON(400,"database transaction error")
		return
	}
	// d.Db.Where("tokenno=?").Find(&compl)
	// c.JSON(200,compl)

	// d.Db.Exec("UPDATE complaints SET status=@status, resolvedate=@resolve associateid=@assid WHERE tokenno=@token",sql.Named("status",data.Status),sql.Named("resolve",Resolvedate),sql.Named("assid",Id),sql.Named("toekn",token))

	var compl models.Complaint
	d.Db.Where("tokenno=?", data.Tokenno).Find(&compl)

	// d.Db.Model(&compl).Select("status","resolvedate","associateid").Updates(map[string]interface{}{"status":data.Status,"resolvedate":Resolvedate,"associateid":Id})
	
	c.JSON(200,compl)
}
