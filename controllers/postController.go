package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"                   // ✅ Keep this
	"github.com/yourusername/go-crud-api/models" // ✅ CORRECT: Import models
)

var posts = []models.Post{
	{ID: 1, Title: "First Post", Content: "This is the first post"},
}

// GET /posts
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

// GET /posts/:id
func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, post := range posts {
		if post.ID == id {
			c.JSON(http.StatusOK, post)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

// POST /posts
func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newPost.ID = len(posts) + 1
	posts = append(posts, newPost)
	c.JSON(http.StatusCreated, newPost)
}

// PUT /posts/:id
func UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedPost models.Post

	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts[i].Title = updatedPost.Title
			posts[i].Content = updatedPost.Content
			c.JSON(http.StatusOK, posts[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

// DELETE /posts/:id
func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}
