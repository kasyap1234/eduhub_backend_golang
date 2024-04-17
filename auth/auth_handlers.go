package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/eduhub_backend_golang/database"
	model "github.com/kasyap1234/eduhub_backend_golang/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

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

type Claims struct {
	UserID string `json:"user_id"`
	Role string `json:"role"`
	jwt.StandardClaims

}
func init() {
    fmt.Println("Initializing database connection and collection...")
    client = database.ConnectDB()
    collection = client.Database("college").Collection("auth")
    fmt.Println("Collection initialized:", collection)
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
        Username string `json:"Username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }
    collection := database.GetMongoClient().Database("college").Collection("auth")
    var user model.User
    filter := bson.M{"Username": credentials.Username}
    err := collection.FindOne(context.Background(), filter).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
        }
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: user.ID.Hex(),
        Role:   user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString, "expires": expirationTime})
}
// c
func hashPassword(password string)(string,error){
	hashedPassword,err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
   if err !=nil {
   return "", err
   }
   return string(hashedPassword),nil
}
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(401, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(*Claims)
        if !ok || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Set user context
        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)

        c.Next()
    }
}

// RoleAuthMiddleware is a middleware for role-based authorization
func RoleAuthMiddleware(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        for _, allowedRole := range roles {
            if role == allowedRole {
                c.Next()
                return
            }
        }
        c.JSON(403, gin.H{"error": "Insufficient permissions"})
        c.Abort()
    }
}
