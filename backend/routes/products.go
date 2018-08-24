package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  "time"
  "log"
)

func ImportProducts(c *gin.Context)  {
  file, err := c.FormFile("file")
  if err != nil {
    c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
    return
  }

  excelFileName:="api/upload/"+"File_"+time.Now().Format("20060102150405")+"_"+file.Filename

  if err := c.SaveUploadedFile(file, excelFileName); err != nil {
    c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
    return
  }

  if readExcel(excelFileName){
    log.Println("End Read file ", excelFileName)
    c.JSON(http.StatusOK, ApiMessage{fmt.Sprintf("File %s uploaded successfully", file.Filename)})
    return
  }

  log.Println("End Read file ",excelFileName)
  c.JSON(http.StatusBadRequest, ApiMessage{"file not save"})

}
