package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kasyap1234/eduhub_backend_golang/auth"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a Gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Allow all origins (CORS)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"content-type"},
		AllowCredentials: true,
	}))

	// Initialize SuperTokens and set up authentication configurations from auth package
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	if err := auth.InitSuperTokens(router, apiBasePath, websiteBasePath); err != nil {
		log.Fatal("Error initializing SuperTokens:", err)
	}
	log.Println("SuperTokens initialized")
	router.POST("/auth/signup")
	router.POST("/auth/signin")

	// Start the server
	router.Run(":8080")
}
