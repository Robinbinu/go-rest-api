package main

import (
	"example.com/db"
	"example.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	//gin server
	server := gin.Default()
	routes.RegisterRoutes(server)
	
	server.Run()
}
