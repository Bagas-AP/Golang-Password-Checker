package main

import (
	"pw/config"
	"pw/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	type PasswordInput struct {
		Pw string `json:"pw"`
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})

	passwordChecker := config.MakePasswordCheckerConfig()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hai Sayang",
		})
	})

	r.POST("/check-password", func(c *gin.Context) {
		var input PasswordInput
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
				"status":  "error",
			})
			return
		}

		res := utils.CheckPassword(input.Pw, passwordChecker)

		if !res {
			c.JSON(400, gin.H{
				"message": "Password is invalid",
				"status":  "error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Password is valid",
			"status":  "success",
		})
	})

	if err := r.Run(":9000"); err != nil {
		panic(err.Error())
	}
}
