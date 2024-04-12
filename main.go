package main

import (
	"log"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"

	// "github.com/kasyap1234/eduhub_backend_golang/auth"
	"github.com/kasyap1234/eduhub_backend_golang/database"
	"github.com/kasyap1234/eduhub_backend_golang/handlers"
)

func main() {
	// Load environment variables from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
    
	router := gin.New()
    

	// Allow all origins (CORS)
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // Allow all origins
	// 	AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	// 	AllowHeaders:     []string{"content-type"},
	// 	AllowCredentials: true,
	// }))
	database.ConnectDB()
	// Initialize SuperTokens and set up authentication configurations from auth package
	// apiBasePath := "/auth"
	// websiteBasePath := "/auth"
	// if err := auth.InitSuperTokens(router, apiBasePath, websiteBasePath); err != nil {
	// 	log.Fatal("Error initializing SuperTokens:", err)
	// }
	// log.Println("SuperTokens initialized")
    log.Default().Println("Server is running on port 8080")
	router.GET("/company/Getallcompanies", handlers.GetAllCompanies)
	router.GET("/company/GetcompanybyID/:id", handlers.GetCompanyByID)
	router.PUT("/company/updateCompany/:id", handlers.UpdateCompany)
	router.POST("/company/createCompany", handlers.AddCompany)
	router.DELETE("/company/DeleteCompany/:id", handlers.DeleteCompany)

	router.Run(":8080")
}
