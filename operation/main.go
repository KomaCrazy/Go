package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dataset struct {
	ID   string `"json":"id"`
	Name string `"json":"name"`
}

var profile = []dataset{
	{ID: "1", Name: "kaw"},
}

func postprofile(c *gin.Context) {
	var newAlbum dataset
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	profile = append(profile, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/profile", postprofile)
	router.Run("localhost:8080")
}
