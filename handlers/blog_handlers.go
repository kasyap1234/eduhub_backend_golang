package handlers 


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

func UpdateBlog(c *gin.Context){
	
}