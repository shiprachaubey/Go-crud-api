// package controllers

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/yourusername/go-crud-api/config"
// 	"github.com/yourusername/go-crud-api/models"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var postCollection *mongo.Collection = config.GetCollection("posts")

// // GET /posts
// func GetPosts(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cursor, err := postCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var posts []models.Post
// 	for cursor.Next(ctx) {
// 		var post models.Post
// 		if err := cursor.Decode(&post); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor decode error"})
// 			return
// 		}
// 		posts = append(posts, post)
// 	}

// 	c.JSON(http.StatusOK, posts)
// }

// // GET /posts/:id
// func GetPost(c *gin.Context) {
// 	idParam := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var post models.Post
// 	err = postCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, post)
// }

// // POST /posts
// func CreatePost(c *gin.Context) {
// 	var post models.Post

// 	if err := c.ShouldBindJSON(&post); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	post.ID = primitive.NewObjectID()

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	_, err := postCollection.InsertOne(ctx, post)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting post"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, post)
// }

// // PUT /posts/:id
// func UpdatePost(c *gin.Context) {
// 	idParam := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	var updatedPost models.Post
// 	if err := c.ShouldBindJSON(&updatedPost); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	update := bson.M{
// 		"$set": bson.M{
// 			"title":   updatedPost.Title,
// 			"content": updatedPost.Content,
// 		},
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	result, err := postCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
// 	if err != nil || result.MatchedCount == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or update failed"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
// }

// // DELETE /posts/:id
// func DeletePost(c *gin.Context) {
// 	idParam := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	result, err := postCollection.DeleteOne(ctx, bson.M{"_id": objID})
// 	if err != nil || result.DeletedCount == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or delete failed"})
// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
//	}
package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-crud-api/config"
	"github.com/yourusername/go-crud-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ✅ Safe way: get collection inside each function
func getPostCollection() *mongo.Collection {
	return config.GetCollection("posts")
}

// GET /posts
func GetPosts(c *gin.Context) {
	postCollection := getPostCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := postCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	for cursor.Next(ctx) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor decode error"})
			return
		}
		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}

// GET /posts/:id
func GetPost(c *gin.Context) {
	postCollection := getPostCollection()

	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var post models.Post
	err = postCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// POST /posts
func CreatePost(c *gin.Context) {
	postCollection := getPostCollection()

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := postCollection.InsertOne(ctx, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// PUT /posts/:id
func UpdatePost(c *gin.Context) {
	postCollection := getPostCollection()

	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":   updatedPost.Title,
			"content": updatedPost.Content,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil || result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
}

// DELETE /posts/:id
func DeletePost(c *gin.Context) {
	postCollection := getPostCollection()

	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
