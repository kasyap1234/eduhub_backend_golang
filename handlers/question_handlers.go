package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/eduhub_backend_golang/database"
)

func GetAllQuestions(c *gin.Context) {
	collection := database.GetMongoClient().Database("college").Collection("questions")

}
func GetQuestionByID(c *gin.Context) {

}

func AddQuestion(c *gin.Context) {

}
func UpdateQuestion(c *gin.Context) {

}
func DeleteQuestion(c *gin.Context) {

}
