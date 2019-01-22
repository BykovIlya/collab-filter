package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/fxsjy/gonn/gonn"
	"fmt"
	"ColabFilter/colab-filter/backend/models"
	"strconv"
)

func GetResult(c *gin.Context) {
	gender, err := strconv.ParseFloat(c.Param("gender"),64)
	age, err := strconv.ParseFloat(c.Param("age"), 64)
	category, err := strconv.ParseFloat(c.Param("category"), 64)
	price,err := strconv.ParseFloat(c.Param("price"), 64)
	if (err != nil) {
		fmt.Println("err with parse string to float64")
	}
	nn := gonn.LoadNN("gonnPerson")
	out := nn.Forward([]float64{gender, age, category, price})
	res := models.GetResult(out)
	fmt.Println("nn result:", res)
	c.JSON(200, res)
}

