package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-crud-api/routers"
)

func main() {
	r := gin.Default()
	routers.RegisterPostRoutes(r)
	r.Run(":8080")
}
