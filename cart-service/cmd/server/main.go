package main

import (
	"cart-service/internal/models"
	"cart-service/internal/routes"
	"cart-service/pkg/config"
	"cart-service/pkg/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	cfg := config.LoadConfig()

	connect := database.Connect(cfg.GetPostgresURL())

	connect.AutoMigrate(&models.Cart{})

	server := gin.Default()
	server.Use(gin.Logger())
	server.SetTrustedProxies([]string{"127.0.0.1"})

	routes := &routes.Routes{
		Db: connect,
	}
	routes.CartRoute(server)

	port, err := strconv.Atoi(cfg.Server.Port)
	if err != nil {
		log.Println(err)
	}
	server.Run(fmt.Sprintf(":%d", port))
}
