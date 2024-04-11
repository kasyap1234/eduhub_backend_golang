package handlers

import (
	"log"
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
	}
	c.JSON(http.StatusOK, companies)
}
func AddCompany(c *gin.Context) error {
	var newCompany model.Company
	if err := c.BindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return err
	}
	collection := database.GetMongoClient().Database("college").Collection("companies")
	err := database.InsertOne(collection, newCompany)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add new company"})
		log.Fatal(err)
	}
	return err

}
func GetCompaniesByID(c *gin.Context) {
	companyId :=c.Param("companyId");
	objID,err :=primitive.ObjectIDFromHex(companyID);
	collection :=databaase.GetMongoClient().Database("college").Collection("companies")

	if err!= nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid company ID"})
		return
	}
	filter :=bson.D{{"_id",objID}};
	company,err :=database.FindOneById(collection, filter);
	if err !=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "failed to fetch company by ID "})
return
	}
	if company ==nil {
		c.JSON(http.StatusNotFound,gin.H{"error": "Company Not found"})
		return
	}
	c.JSON(http.StatusOK,company);


}
func UpdateCompany(c *gin.Context) {

}
func DeleteCompany(c *gin.Context) {

}
