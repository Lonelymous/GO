package main

import "github.com/gin-gonic/gin"

var PORT string = ":8080"

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"v1": "login",
			})
		})
		v1.GET("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"v1": "submit",
			})
		})
		v1.GET("/read", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"v1": "read",
			})
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"v2": "login",
			})
		})
		v2.GET("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"v2": "submit",
			})
		})
		v2.GET("/read", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"v2": "read",
			})
		})
	}

	router.Run(PORT)
}