package main

import (
	"log"

	"github.com/VJ-Vijay77/customerServiceMiniProject/initializers"
	"github.com/VJ-Vijay77/customerServiceMiniProject/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	Init()
	r := gin.Default()
	routes.Routes(r)
	if err := r.Run(":8080") ;err != nil{
		log.Fatalln(err)
	}

}