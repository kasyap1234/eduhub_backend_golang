package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/kasyap1234/eduhub_backend_golang/database"
    model "github.com/kasyap1234/eduhub_backend_golang/models"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "gopkg.in/mgo.v2/bson"
)

// GetAllCompanies godoc
// @Summary Get all companies
// @Description Get details of all companies
// @Tags companies
// @Produce json
// @Success 200 {object}
// @Failure 500 {object} gin.H{"error": string}
// @Router /companies [get]
func GetAllCompanies(c *gin.Context) {
    collection := database.GetMongoClient().Database("college").Collection("companies")
    companies, err := database.FindAll(collection)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch all the companies"})
        return
    }
    c.JSON(http.StatusOK, companies)
}

// AddCompany godoc
// @Summary Add a new company
// @Description Add a new company to the database
// @Tags companies
// @Accept json
// @Produce json
// @Param company body model.Company true "New Company"
// @Success 200 {object} gin.H{"message": string}
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /companies [post]
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

// GetCompanyByID godoc
// @Summary Get a company by ID
// @Description Get details of a specific company by ID
// @Tags companies
// @Produce json
// @Param companyId path string true "Company ID"
// @Success 200 {object} model.Company
// @Failure 400 {object} gin.H{"error": string}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /companies/{companyId} [get]
func GetCompanyByID(c *gin.Context) {
    companyID := c.Param("companyId")
    objID, err := primitive.ObjectIDFromHex(companyID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company ID"})
        return
    }
    collection := database.GetMongoClient().Database("college").Collection("companies")
    filter := bson.D{{"_id", objID}}
    company, err := database.FindOneById(collection, filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch company by ID"})
        return
    }
    if company == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
        return
    }
    c.JSON(http.StatusOK, company)
}

// UpdateCompany godoc
// @Summary Update a company
// @Description Update details of an existing company
// @Tags companies
// @Accept json
// @Produce json
// @Param companyId path string true "Company ID"
// @Param company body model.Company true "Updated Company"
// @Success 200 {object} gin.H{"message": string}
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /companies/{companyId} [put]
func UpdateCompany(c *gin.Context) {
    var updatedCompany model.Company
    if err := c.BindJSON(&updatedCompany); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
    companyID := c.Param("companyId")
    objID, err := primitive.ObjectIDFromHex(companyID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company ID"})
        return
    }
    collection := database.GetMongoClient().Database("college").Collection("companies")
    filter := bson.D{{"_id", objID}}
    update := bson.D{{"$set", bson.D{{"name", updatedCompany.Name}, {"url", updatedCompany.Url}}}}
    err = database.UpdateOne(collection, filter, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

// DeleteCompany godoc
// @Summary Delete a company
// @Description Delete a company by ID
// @Tags companies
// @Produce json
// @Param companyId path string true "Company ID"
// @Success 200 {object} gin.H{"message": string}
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /companies/{companyId} [delete]
func DeleteCompany(c *gin.Context) {
    companyID := c.Param("companyId")
    objID, err := primitive.ObjectIDFromHex(companyID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company ID"})
        return
    }
    collection := database.GetMongoClient().Database("college").Collection("companies")
    filter := bson.D{{"_id", objID}}
    err = database.DeleteOne(collection, filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
