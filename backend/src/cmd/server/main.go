package main

import (
	"log"

	"noise-complaint-backend/src/config"
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	err := database.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	gin.SetMode(cfg.Server.Mode)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-Operator")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.SetupRoutes(r)

	log.Printf("Server starting on port %s...", cfg.Server.Port)
	err = r.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
