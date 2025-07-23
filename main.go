// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/yourusername/go-crud-api/config"
// 	"github.com/yourusername/go-crud-api/routers"
// )

// func main() {
// 	app := gin.Default()

// 	// Connect to MongoDB
// 	config.ConnectDB()

// 	// Register routes
// 	routers.RegisterPostRoutes(app)

// 	app.Run(":8080")
// }

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yourusername/go-crud-api/config"
	"github.com/yourusername/go-crud-api/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Connect to DB

	config.ConnectDB()

	r := gin.Default()
	routers.RegisterPostRoutes(r)
	r.Run(":8080")
}
