package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type data_list struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data_set = []data_list{
	{ID: "1", Name: "kaw"},
	{ID: "2", Name: "gank"},
	{ID: "3", Name: "jan"},
}

func get_data(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data_set)
}

func post_data(c *gin.Context) {
	var newdata data_list
	if err := c.BindJSON(&newdata); err != nil {
		return
	}
	data_set = append(data_set, newdata)
	c.IndentedJSON(http.StatusCreated, newdata)
}

func get_id(c *gin.Context) {
	id := c.Param("id")
	for _, i := range data_set {
		if id == i.ID {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "fail"})

}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/data", get_data)
	router.POST("/data", post_data)
	router.POST("/data/:id", get_id)
	router.Use(cors.Default())
	router.Run("localhost:8080")
}
