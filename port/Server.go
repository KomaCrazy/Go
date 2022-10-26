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
	ID       string `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func login(c *gin.Context) {
	var data data_set
	if err := c.BindJSON(&data); err != nil {
		return
	}
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	cmd := `select * from db where user="` + data.User + `" and password="` + data.Password + `";`
	rows, err := db.Query(cmd)
	for rows.Next() {
		var id string
		var user string
		var password string
		var email string

		err = rows.Scan(&id, &user, &password, &email)
		if err != nil {
			log.Fatal(err)
		}
		data := &data_set{ID: id, User: user, Password: password, Email: email}
		c.IndentedJSON(http.StatusOK, data)
	}
}

func register(c *gin.Context) {
	var newdata data_set
	if err := c.BindJSON(&newdata); err != nil {
		return
	}

}

func main() {
	route := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	route.Use(cors.New(config))

	route.GET("/login", login)
	route.GET("/register", register)
	route.Run("localhost:8000")
}
