package main

import (
	"errors"
	"fmt"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var API = "/api/v1"

var users = []User{
	{
		Email:    "example@example.com",
		Password: "1234567890",
	},
}

func findUserByEmail(email string) (*User, error) {
	for i, u := range users {
		if u.Email == email {
			return &users[i], nil
		}
	}
	return nil, errors.New("User not found")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := findUserByEmail(email)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func loginUser(c *gin.Context) {
	email, hasEmail := c.GetQuery("email")
	password, hasPass := c.GetQuery("password")
	fmt.Printf("%t", hasEmail)
	fmt.Printf("%t", hasPass)
	if !hasEmail {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing email parameter"})
		return
	}

	if !hasPass {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing password parameter"})
		return
	}

	user, err := findUserByEmail(email)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Email not found"})
		return
	}

	if user.Password != password {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Password did not match"})
		return
	}

	c.IndentedJSON(http.StatusFound, gin.H{"message": "Login successful"})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
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

	r.GET("/api/v1/users/:email", getUserByEmail)

	r.POST("/api/v1/users", createUser)

	r.GET("/api/v1/users/login", loginUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
