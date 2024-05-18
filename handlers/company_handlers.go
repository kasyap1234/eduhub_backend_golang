package handlers

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/eduhub_backend_golang/database"
	model "github.com/kasyap1234/eduhub_backend_golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	
)

func GetAllCompanies(c *gin.Context) {
	collection := database.GetMongoClient().Database("college").Collection("companies")
	companies, err := database.FindAll(collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch all the companies"})
		return // Added return statement
	}
	c.JSON(http.StatusOK, companies)
}

// func AddCompany(c *gin.Context) {
// 	var newCompany model.Company
// 	if err := c.BindJSON(&newCompany); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
// 		return // Added return statement
// 	}
// 	collection := database.GetMongoClient().Database("college").Collection("companies")
// 	err := database.InsertOne(collection, newCompany)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add new company"})
// 		log.Fatal(err)
// 		return // Added return statement
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Company added successfully"})
// }

func AddCompany(c *gin.Context) {
    var newCompany model.Company
    if err := c.BindJSON(&newCompany); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }
    collection := database.GetMongoClient().Database("college").Collection("companies")
    err := database.InsertOne(collection, newCompany)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add new company"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Company added successfully"})
}
func GetCompanyByID(c *gin.Context) {
	companyID := c.Param("companyId") // Changed to "companyId"
	objID, err := primitive.ObjectIDFromHex(companyID)
	collection := database.GetMongoClient().Database("college").Collection("companies")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company ID"})
		return // Added return statement
	}
	filter := bson.D{{"companyId", objID}}
	company, err := database.FindOneById(collection, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch company by ID"})
		return // Added return statement
	}
	if company == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return // Added return statement
	}
	c.JSON(http.StatusOK, company)
}

func UpdateCompany(c *gin.Context) {
	var updatedCompany model.Company
	if err := c.BindJSON(&updatedCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return // Added return statement
	}
	companyID := c.Param("companyId")
	objID, err := primitive.ObjectIDFromHex(companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company ID"})
		return // Added return statement
	}
	filter := bson.D{{"companyId", objID}}
	update := bson.D{{"$set", bson.D{{"name", updatedCompany.Name}, {"Url", updatedCompany.Url}}}}
	collection := database.GetMongoClient().Database("college").Collection("companies")
	database.UpdateOne(collection, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return // Added return statement
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

func DeleteCompany(c *gin.Context) {
	companyID := c.Param("companyId")
	objID, err := primitive.ObjectIDFromHex(companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company ID"})
		return // Added return statement
	}
	filter := bson.D{{"companyId", objID}}
	collection := database.GetMongoClient().Database("college").Collection("companies")
	err = database.DeleteOne(collection, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return // Added return statement
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
