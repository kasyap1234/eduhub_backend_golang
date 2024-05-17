package main

import (
	"log"
	
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
    // _ "github.com/kasyap1234/eduhub_backend_golang/docs"
	"github.com/kasyap1234/eduhub_backend_golang/auth"
	// "github.com/kasyap1234/eduhub_backend_golang/auth"
	"github.com/kasyap1234/eduhub_backend_golang/database"
	"github.com/kasyap1234/eduhub_backend_golang/handlers"
	 swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)
// @title EduHub API
// @version 1.0
// @description This is a sample server for EduHub.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Load environment variables from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
    
	router := gin.New()
    


    // Enable CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Allow all origins
        AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))
	database.ConnectDB()
	// Initialize SuperTokens and set up authentication configurations from auth package
	// apiBasePath := "/auth"
	// websiteBasePath := "/auth"
	// if err := auth.InitSuperTokens(router, apiBasePath, websiteBasePath); err != nil {
	// 	log.Fatal("Error initializing SuperTokens:", err)
	// }
	// log.Println("SuperTokens initialized")
    log.Default().Println("Server is running on port 8080")
	r1 := router.Group("/company") 
	{
	r1.GET("/Getallcompanies", handlers.GetAllCompanies)
	r1.GET("/GetcompanybyID/:id", handlers.GetCompanyByID)
	r1.PUT("/updateCompany/:id", handlers.UpdateCompany)
	r1.POST("/createCompany", handlers.AddCompany)
	r1.DELETE("/DeleteCompany/:id", handlers.DeleteCompany)
	}
	r2 :=router.Group("/question")
	{
		r2.GET("/Getallquestions", handlers.GetAllQuestions)
		r2.GET("/GetquestionbyID/:id", handlers.GetQuestionByID)
		r2.PUT("/updateQuestion/:id", handlers.UpdateQuestion)
		r2.POST("/createQuestion", handlers.AddQuestion)
		r2.DELETE("/DeleteQuestion/:id", handlers.DeleteQuestion)
        
	}
	r3 :=router.Group("/auth")
	{

		r3.POST("/login", auth.LoginUser)
		r3.POST("/signup", auth.RegisterUser)
		
	}
 r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
