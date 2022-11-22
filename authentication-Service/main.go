package main

import (
	"log"

	"github.com/VJ-Vijay77/authServiceMiniProject/initializers"
	"github.com/VJ-Vijay77/authServiceMiniProject/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.ConnectToDB()
}
func main() {
	//initializing the database at startup
	Init()
	r := gin.Default()
	routes.Routes(r)
	if err := r.Run(":8081"); err != nil {
		log.Fatalln(err)
	}
}
