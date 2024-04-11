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

func GetAllQuestions(c *gin.Context) {
	collection := database.GetMongoClient().Database("college").Collection("questions")
	questions, err := database.FindAll(collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
		return
	}
	c.JSON(http.StatusOK, questions)
}
func GetQuestionByID(c *gin.Context) {
	questionID := c.Param("questionID")
	objID, err := primitive.ObjectIDFromHex(questionID)
	collection := database.GetMongoClient().Database("college").Collection("questions")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect  request body  "})
		return
	}
	filter := bson.D{{"questionID",objID}}
	question, err := database.FindOneById(collection, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch question by ID "})
		return
	}
	if question == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}
	return

}

func AddQuestion(c *gin.Context) {
	var newQuestion model.Question
	if err := c.BindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	collection := database.GetMongoClient().Database("college").Collection("questions")
	err := database.InsertOne(collection, newQuestion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add question"})
		log.Fatal(err)

	}
	return

}
func UpdateQuestion(c *gin.Context) {
	var updatedQuestion model.Question
	if err := c.BindJSON(&updatedQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	questionID := c.Param("questionID")
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Question ID "})
	}

	filter := bson.D{{"questionId", objID}}
	update := bson.D{{"$set", bson.D{{"Title", updatedQuestion.Title}, {"Text", updatedQuestion.Text}, {"Answer", updatedQuestion.Answer}}}}
     collection := database.GetMongoClient().Database("college").Collection("questions")
	  database.UpdateOne(collection, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update question"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "question updated successfully "})

}

func DeleteQuestion(c *gin.Context) {
	questionID := c.Param("questionID")
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id "})
		return
	}
	filter := bson.D{{"questionId", objID}}
	collection := database.GetMongoClient().Database("college").Collection("questions")
	err = database.DeleteOne(collection, filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete question"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted question "})

}
