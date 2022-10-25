package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type data_set struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data = []data_set{
	{ID: "1", Name: "kaw"},
}

func get_data(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}

func post_data(c *gin.Context) {
	var newdata data_set
	if err := c.BindJSON(&newdata); err != nil {
		return
	}
	data = append(data, newdata)
	c.IndentedJSON(http.StatusCreated, newdata)

}

func main() {
	route := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	route.Use(cors.New(config))
	route.GET("/", get_data)
	route.POST("/", post_data)
	route.Run("localhost:8080")
}
