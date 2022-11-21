package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Session struct {
	Userid int
	Expiry time.Time
}

var Sessions = map[string]Session{}

func IsExpired(c *gin.Context) bool{
	cookies,err := c.Request.Cookie("session_token")
	if err != nil {
		
		return false
	}
	sesToken := cookies.Value

	userSession,exists := Sessions[sesToken]
	if !exists {
		
		return false
	}
	if !time.Now().Before(userSession.Expiry) {
		return false
	}
	
	return true
}


func CreateSession(ID int,c *gin.Context) string{
	
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(30 *time.Second)

	Sessions[sessionToken] = Session{
		Userid: ID,
		Expiry: expiresAt,
	}

	http.SetCookie(c.Writer,&http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: expiresAt.Add(2 * time.Minute),
	})
	return sessionToken
}
