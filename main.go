package main

import (
	"fmt"
	"identity-service/model"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if strings.ToLower(os.Getenv("GIN_MODE")) != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", Login)

	r.Run(fmt.Sprintf(":%s", os.Getenv("REST_PORT")))
}

// Login ...
func Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided")
		return
	}

	c.JSON(http.StatusOK, "valid JSON!")
}
