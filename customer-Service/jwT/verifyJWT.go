package jwT

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJWT(c *gin.Context) bool {
	tokenstring, err := c.Request.Cookie("jwt")
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"status": "Couldnt get coookie",
	// 		"error":  err,
	// 	})
	// 	return false
	// }

	if err != nil {
		return false
	}

	token, err := jwt.Parse(tokenstring.Value, func(t *jwt.Token) (interface{}, error) {
		return SampleSecretKey, nil
	})

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error":  err,
	// 		"status": "Token is not valid",
	// 	})
	// 	return false
	// }
	if err != nil {
		return false
	}

	if !token.Valid {
		// c.JSON(401, gin.H{
		// 	"status": "token is not valid",
		// })
		return false
	}

	// c.JSON(200, gin.H{
	// 	"status": "Token is valid",
	// })
	return true
}


func GetIDofUser(c *gin.Context) uint{
	cookie,err := c.Request.Cookie("jwt")
	if err != nil {
		c.JSON(400,"error getting cookies")
		return 0
	}
	token := cookie.Value
	claims := &Claims{}

	tkn,err := jwt.ParseWithClaims(token,claims,func(t *jwt.Token) (interface{}, error) {
		return SampleSecretKey,nil
	})
	if err != nil {
		c.JSON(401,"Token invalid or expired, please login again")
		return 0
	}
	if !tkn.Valid {
		c.JSON(401,"1:Token invalid or expired, please login again")
		return 0
	}
	Id := claims.UserID
	return Id
}