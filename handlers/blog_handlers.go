package handlers 
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/eduhub_backend_golang/database"
	model "github.com/kasyap1234/eduhub_backend_golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)
func getAllBlogs(c *gin.Context){
	collection :=database.GetMongoClient().Database("college").Collection("blogs"); 
	blogs,err := database.FindAll(collection); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return 
	}
	c.JSON(http.StatusOK, blogs)

}
func addBlog(c *gin.Context){
	var newBlog model.Blog 
	if err := c.BindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	collection := database.GetMongoClient().Database("college").Collection("blogs")
	err := database.InsertOne(collection, newBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog added successfully"})
}
func getBlogByID(c *gin.Context){
 blogID :=c.Param("ID")
 objID,err :=primitive.ObjectIDFromHex(blogID)
 if err !=nil {
	 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
	 return

 }
 collection :=database.GetMongoClient().Database("college").Collection("blogs")
 filter :=bson.D{{"ID",objID}}
 blog,err :=database.FindOneById(collection,filter); 
 if err !=nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog"})
	return 
 }
 c.JSON(http.StatusOK, blog)


}
func deleteBlogByID(c *gin.Context){
	blogId :=c.Param("ID")
	objId,err :=primitive.ObjectIDFromHex(blogId)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	collection :=database.GetMongoClient().Database("college").Collection("blogs")
	filter :=bson.D{{"ID",objId}}
	err =database.DeleteOne(collection,filter); 
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})

}
func UpdateBlog(c *gin.Context){
    blogId :=c.Param("ID")
	objId,err :=primitive.ObjectIDFromHex(blogId)
	if err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	var updatedBlog model.Blog
	if err := c.BindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	collection :=database.GetMongoClient().Database("college").Collection("blogs")
	filter :=bson.D{{"ID",objId}}
	update :=bson.D{{"$set",bson.D{{"title",updatedBlog.Title},{"text",updatedBlog.Text},{"author",updatedBlog.Author}}}}
	database.UpdateOne(collection,filter,update);
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return 
	}

}