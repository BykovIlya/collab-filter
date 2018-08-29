package main

import (
	"path/filepath"
	"os"
  "fmt"
  "github.com/gin-gonic/gin"
  "ColabFilter/colab-filter/backend/routes"
)

func main() {
	fmt.Println("Hello!")
	CreateDirsForFiles()

	router := gin.Default()
  router.Use(CORSMiddleware())
	router.Static("/api/c/tmp", "./api/c/tmp")
  router.POST("/import", routes.ImportEvents)
	router.GET("/events",routes.GetEvents)
  router.GET("/users", routes.GetUsers)
	router.GET("/products", routes.GetProducts)
  router.GET("/recommendations", routes.GetRecommends)
	router.POST("/recommendations/personal", routes.GetPerson)
  router.Run(":5000")
}

func CreateDirsForFiles() {
	tmp := filepath.Join(".", "api/c/tmp")
	os.MkdirAll(tmp, os.ModePerm)

	upload := filepath.Join(".", "api/upload")
	os.MkdirAll(upload, os.ModePerm)
}
func CORSMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Max-Age", "86400")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length,X-Requested-With, Accept-Encoding, X-CSRF-Token, Authorization")
    c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

    if c.Request.Method == "OPTIONS" {
      fmt.Println("OPTIONS")
      c.AbortWithStatus(200)
    } else {
      c.Next()
    }
  }
}
