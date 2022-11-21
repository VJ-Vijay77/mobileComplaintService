package main

import (
	"log"

	"github.com/VJ-Vijay77/associateServiceMiniProject/initializers"
	"github.com/VJ-Vijay77/associateServiceMiniProject/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.ConnectToDB()
}
func main() {
	Init()
	r := gin.Default()
	routes.Routes(r)
	if err := r.Run(":8082"); err != nil {
		log.Fatalln(err)
	}
}
