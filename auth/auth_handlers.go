// package auth
package auth

import (
	"github.com/gin-gonic/gin"
	model "github.com/kasyap1234/eduhub_backend_golang/models"
	"go.mongodb.org/mongo-driver/mongo"
)
var (
	jwtKey =[]byte("your_secret_key")
	mongoURI = "mongodb://localhost:27017"
	dbName="college"
	collectionName="users"
	client *mongo.Client
	collection *mongo.Collection

)
func init(){
	collection :=client.Database(dbName).Collection(collectionName)

}
func RegisterUser(c *gin.Context){
	var newUser model.User
if err := c.ShouldBindJSON()

}
