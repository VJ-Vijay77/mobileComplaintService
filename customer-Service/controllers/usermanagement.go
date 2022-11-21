package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/VJ-Vijay77/customerServiceMiniProject/initializers"
	"github.com/VJ-Vijay77/customerServiceMiniProject/jwT"
	"github.com/VJ-Vijay77/customerServiceMiniProject/models"
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

func (d *Data) Home(c *gin.Context) {
	ok := jwT.VerifyJWT(c)
	if !ok {
		c.JSON(200, "Welcome to Home,Please login")
		return
	}
	c.JSON(200, "Welcome to Home!")
}

func (d *Data) AddUser(c *gin.Context) {
	ok := jwT.VerifyJWT(c)
	if !ok {
		c.JSON(401, "Please login with your credentials..!")
		return
	}
	var user models.User
	var check models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, "Couldnt get data from!")
		return
	}

	d.Db.Where("phone=?", user.Phone).Find(&check)

	if check.Phone == user.Phone {
		c.JSON(400, gin.H{
			"message":  "The user is already exist in the database",
			"username": check.Firstname,
			"hint":     "Try using different phone number",
		})
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	res := d.Db.Create(&user)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"message": "Coudlnt create the database entry",
			"error":   res.Error,
		})
	}
	c.JSON(200, gin.H{
		"status":        "Created Successfuly",
		"rows affected": res.RowsAffected,
	})

}

func (d *Data) Login(c *gin.Context) {
	ok := jwT.VerifyJWT(c)
	if ok {
		c.JSON(200, "you are already logged in ..!")
		return
	}
	var user models.User
	// var result models.User
	// var check models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, "Couldnt bind json")
		return
	}

	// resp,err := http.G
	data, _ := json.Marshal(user)
	resp, _ := http.NewRequest("POST", "http://authentication:8081/authenticate", bytes.NewBuffer(data))
	client := &http.Client{}
	res, _ := client.Do(resp)

	status := res.StatusCode
	reader := res.Body
	cont, _ := io.ReadAll(reader)
	var k int
	json.Unmarshal(cont, &k)
	switch status {
	case 200:
		var userData models.User

		d.Db.Where("id=?", k).Find(&userData)

		token, err := jwT.JWT(userData.ID)
		if err != nil {
			c.String(400, "fails to issue JWT \n%s", err)
			return
		}
		// c.Request.Header.Set("Token",token)
		// IDofUser,err := jwt.ExtractClaims(c)
		// if err != nil{
		// 	c.String(400,"%s",err)
		// 	return
		// }
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "jwt",
			Value:   token,
			Expires: time.Now().Add(2 * time.Minute),
		})

		c.JSON(200, gin.H{
			"status":               "Login successful",
			"status code recieved": status,
			"jwt token":            token,
		})
	case 401:
		c.JSON(401, gin.H{
			"status": "Wrong credentials",
			"hint":   "phone number or password is wrong",
		})
	default:
		c.JSON(500, "some internal error occured")
	}
}

// resp, _ := http.NewRequest("GET", "http://authentication:8081/ping", nil)
// client := &http.Client{}
// res, _ := client.Do(resp)
// reader := res.Body
// contentlength := res.ContentLength
// contentType := res.Header.Get("Content-Type")
// c.DataFromReader(200, contentlength, contentType, reader, nil)

// data,_ := json.Marshal(user)

// resp,err := http.Post("http://authentication:8081/getuser","application/json",bytes.NewBuffer(data))
// if err != nil {
// 	c.JSON(400,gin.H{
// 		"error":err,
// 	})
// }
// resp.Body.Close()
// content,_ := ioutil.ReadAll(resp.Body)
// json.Unmarshal(content,&result)
// c.JSON(200,gin.H{
// 	"content":string(content),
// })

/*

	d.Db.Find(&check).Where("phone=?", user.Phone)
	if user.Phone == check.Phone && user.Password == check.Password {
		c.JSON(200, gin.H{
			"Status":  "Login Successful",
			"Message": "Welcome " + check.Firstname + " " + check.Lastname,
		})
		// c.String(200,fmt.Sprintf("\nWelcome %s %s",check.Firstname,check.Lastname))
		return
	}
	c.JSON(404, gin.H{
		"status": "Login failed",
		"hint":   "Re-check phone or password",
	})
*/

func (d *Data) Dashboard(c *gin.Context) {
	ok := jwT.VerifyJWT(c)
	if !ok {
		c.JSON(401, "Please login with your credentials..!")
		return
	}
	var Data []models.Complaint
	ID := jwT.GetIDofUser(c)
	
		d.Db.Where("user_id=?", ID).Find(&Data)
		c.JSON(200, Data)
}

func (d *Data) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "jwt",
		Value:  "",
		MaxAge: -1,
	})
	// http.SetCookie(c.Writer, &http.Cookie{
	// 	Name:   "userdata",
	// 	Value:  "",
	// 	MaxAge: -1,
	// })
	c.JSON(200, "Logout successful")
}
