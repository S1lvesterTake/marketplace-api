package api

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/S1lvesterTake/marketplace-api/infrastructure/persistance/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitRoute() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "*"},
		ExposeHeaders:    []string{"Accept", "Content-Length", "Content-Type", "Authorization", "Accept:Encoding"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// v1 := router.Group("/api/v1")

	// Health check
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm well",
		})
	})

	db.DBInit()

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal(fmt.Sprintf("PORT must be set [%s]", port))
	}

	router.Run(":" + port)
}
