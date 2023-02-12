package main

import (
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`

}

var users = []User{
	{
		Email:    "example@example.com",
		Password: "1234567890",
	},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {

	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	r.GET("/api/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.GET("/api/v1/users", getUsers)

	r.POST("/api/v1/users", createUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
