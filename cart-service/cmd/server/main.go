package main

import (
	"cart-service/internal/cache"
	"cart-service/internal/routes"
	"cart-service/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	config.LoadConfig()

	cache.ConnectRedis()

	server := gin.Default()
	server.Use(gin.Logger())
	server.SetTrustedProxies([]string{"127.0.0.1"})

	routes := &routes.Routes{}
	routes.CartRoute(server)

	port, err := strconv.Atoi(config.Server.Port)
	if err != nil {
		log.Println(err)
	}
	server.Run(fmt.Sprintf(":%d", port))
}
