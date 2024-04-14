// package auth
package auth

import (
	"context"
	"fmt"
	"net/http"

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
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = hashedPassword

	collection := database.GetMongoClient().Database("college").Collection("auth")
	err = database.InsertOne(collection, newUser)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})

}
func LoginUser(c *gin.Context) {
	var credentials struct {
		Username string `json: "Username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	var user User
	collection := database.GetMongoClient().Database("college").Collection("auth")
	filter := bson.M{"Username": credentials.Username}

	results, error := database.FindOneById(collection, filter)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
   return
	}
	 err :=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte (credentials.Password)){
		c.JSON(401,gin.H{"errror": "invalid credentials"})

		return
	}
     expirationTime :=time.Now().Add(24* time.Hour)
 claims :=&Claims{
UserId : user.ID.Hex(),
Role: user.Role,
StandardClaims: jwt.StandardClaims{
	ExpiresAt: expirationTime.Unix(),

}
token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims);
tokenString, err:=token.SignedString(jwtkey);
if err !=nil {
	c.JSON(500,gin.H{"error": "Failed to generate token"})
	return
}
c.JSON(200,gin.H{"token": tokenString, "expires": expirationTime})
 }

func hashPassword(password string)(string,error){
	hashedPassword,err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
   if err !=nil {
   return "", err
   }
   return string(hashedPassword),nil
}
func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString :=c.GetHeader("Authorization")
		if tokenString == ""{
			c.JSON(401,gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		token, err :=jwt.ParseWithClaims(tokenString,&Claims{},)
	}
}
