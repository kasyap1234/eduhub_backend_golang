package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kasyap1234/eduhub_backend_golang/auth" // Import your auth package
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Create a Gin router
    router := gin.Default()

    // Allow CORS for your frontend domain
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{os.Getenv("WEBSITE_DOMAIN")},
        AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
        AllowHeaders:     []string{"content-type"},
        AllowCredentials: true,
    }))

    // Initialize SuperTokens and set up authentication configurations from auth package
    if err := auth.InitSuperTokens(router); err != nil {
        log.Fatal("Error initializing SuperTokens:", err)
    }

    // Start the server
    router.Run(":8080")
}
