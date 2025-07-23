package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-crud-api/controllers"
)

func RegisterPostRoutes(router *gin.Engine) {
	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/:id", controllers.GetPost)
		postRoutes.POST("/", controllers.CreatePost)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}
}
