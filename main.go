package main

import (
	"log"
	
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
"github.com/gin-contrib/cors"
	"github.com/kasyap1234/eduhub_backend_golang/auth"
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

	router.Run(":8080")
}
