// package auth
package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/eduhub_backend_golang/database"
	model "github.com/kasyap1234/eduhub_backend_golang/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var (
	jwtKey         = []byte("your_secret_key")
	mongoURI       = "mongodb://localhost:27017"
	dbName         = "college"
	collectionName = "users"
	client         *mongo.Client
	collection     *mongo.Collection
)

func init() {
	collection := client.Database(dbName).Collection(collectionName)
	fmt.Print(collection)
}
func RegisterUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	// check if username already exists in the database ;
	var existingUser model.User
	err := collection.FindOne(context.Background(), bson.M{"username": newUser.Username}).Decode(&existingUser)
	if err == nil {
		c.JSON(400, gin.H{"error": "username already exists"})
		return
	}
	// hash password ;
	collection := database.GetMongoClient().Database("college").Collection("auth")
	err = database.InsertOne(collection, newUser)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})

}
func LoginUser(c *gin.Context) {

}
