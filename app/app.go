package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type data_set struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

var data = []data_set{}

func found(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from db;")
	for rows.Next() {
		var id string
		var name string
		var age string
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		newdata := &data_set{ID: id, Name: name, Age: age}
		data = append(data, *newdata)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func insert(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	route := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	route.Use(cors.New(config))

	route.GET("/", found)

	route.Run("localhost:8000")
}
