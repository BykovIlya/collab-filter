package main

import (
	"path/filepath"
	"os"
  "fmt"
  "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello!")
	CreateDirsForFiles()

	router := gin.Default()
	router.Static("/api/c/tmp", "./api/c/tmp")
	//router.Use(CORSMiddleware())
  router.Run(":5000")
}

func CreateDirsForFiles() {
	tmp := filepath.Join(".", "api/c/tmp")
	os.MkdirAll(tmp, os.ModePerm)

	upload := filepath.Join(".", "api/upload")
	os.MkdirAll(upload, os.ModePerm)
}
