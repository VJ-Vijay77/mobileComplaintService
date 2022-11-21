package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/customerServiceMiniProject/jwT"
	"github.com/VJ-Vijay77/customerServiceMiniProject/models"
	"github.com/gin-gonic/gin"
)

func (d *Data) Complaints(c *gin.Context) {
	ok := jwT.VerifyJWT(c)
	if !ok {
		c.JSON(401, "Please login with your credentials..!")
		return
	}
	var complaints models.Complaint

	if err := c.ShouldBindJSON(&complaints); err != nil {
		c.JSON(400, "Failed to bind JSON")
		return
	}
	ID := jwT.GetIDofUser(c)
	complaints.UserID = ID
	complaints.Complaintdate = time.Now().Format("01-02-2006")
	t := d.Db.Create(&complaints)
	if t.Error != nil {
		c.JSON(401, "Failed to register complaint, try again")
		return
	}
	c.JSON(200, "Complaint Registered Successfully")
}

func (d *Data) GetStatus(c *gin.Context) {
	token := c.Param("token")
	var status models.Complaint

	t := d.Db.Select("tokenno", "user_id", "complaints", "complaintdate", "status", "associateid").Where("tokenno=?", token).Find(&status)
	if t.Error != nil {
		c.JSON(400, "Couldnt find the record")
		return
	}

	c.JSON(200, gin.H{
		"Token No":       status.Tokenno,
		"User ID":        status.UserID,
		"Complaint":      status.Complaints,
		"Complaint Date": status.Complaintdate,
		"Status":         status.Status,
		"Associate ID":   status.Associateid,
	})
}

func (d *Data) ComplaintStatus(c *gin.Context) {
	token := c.Param("token")
	var stat models.Complaint
	tkn, _ := strconv.Atoi(token)
	url := fmt.Sprintf("http://associate:8082/status/%d", tkn)
	res, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	response, _ := client.Do(res)

	status := response.StatusCode
	reader := response.Body
	cont, _ := io.ReadAll(reader)
	json.Unmarshal(cont,&stat)
	switch status {
	case 200:
		c.JSON(200,stat)
		return
	case 400:
		c.JSON(400,"Invalid token no or some internal error")
		return
	default:
		c.JSON(400,"Some problem occured try again")		
	}
}
