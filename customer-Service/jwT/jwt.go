package jwT

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var SampleSecretKey = []byte("SecretYouShouldHide")

type Claims struct {
	UserID uint
	jwt.RegisteredClaims
}

// func IssueJwt(ID uint, c *gin.Context) (string, error) {
// 	T := jwt.NewNumericDate(time.Now().Add(time.Minute * 3))

// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
// 		Issuer:    strconv.Itoa(int(ID)),
// 		ExpiresAt: T,
// 	})

// 	token, err := claims.SignedString([]byte(SampleSecretKey))
// 	if err != nil {
// 		return "", err
// 	}

// 	// cookie := &http.Cookie{
// 	// 	Name: "jwt",
// 	// 	Value: token,
// 	// 	Expires: time.Now().Add(time.Minute *5),
// 	// }
// 	val := strconv.Itoa(int(ID))
// 	// c.SetCookie("jwt", token, 5, "", "", false, true)
// 	// c.SetCookie("userdata",val,5,"","",false,true)
// 	http.SetCookie(c.Writer, &http.Cookie{
// 		Name:    "jwt",
// 		Value:   token,
// 		Expires: time.Now().Add(time.Minute * 1),
// 	})
// 	http.SetCookie(c.Writer, &http.Cookie{
// 		Name:    "userdata",
// 		Value:   val,
// 		Expires: time.Now().Add(time.Minute * 1),
// 	})
// 	return token, nil
// }

// func JWTtokenGenerate(ID uint) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["exp"] = time.Now().Add(1 * time.Minute)
// 	claims["authorizes"] = true
// 	claims["user"] = ID

// 	tokenString, err := token.SignedString(SampleSecretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func JWT(ID uint) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Minute * 1))
	// expireTime := time.Now().Add(1 * time.Minute)
	claims := Claims{
		UserID: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func VerifyJWT(c *gin.Context) (error, bool) {
// 	req := c.Request
// 	var check bool
// 	if req.Header["Token"] != nil {

// 		token, err := jwt.Parse(req.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
// 			_, ok := t.Method.(*jwt.SigningMethodHMAC)
// 			if !ok {
// 				check = false
// 				return "", errors.New("signIn method in failed")
// 			}

// 			return "", nil
// 		})

// 		if err != nil {
// 			return err, false
// 		}

// 		if token.Valid {
// 			check = true
// 			return nil, check
// 		} else {
// 			check = false
// 			return nil, check
// 		}
// 	}
// 	check = false
// 	return nil, check
// }

func ExtractClaims(c *gin.Context) (string, error) {
	if c.Request.Header["Token"] != nil {
		tokenString := c.Request.Header["Token"][0]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return nil, nil
		})
		if err != nil {
			return "", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userID := claims["user"].(string)
			return userID, nil
		}
		//  return "",errors.New("Cant get data from claims")

	}
	return "", errors.New("cant get data from claims")
}
