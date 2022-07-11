package main

import "github.com/gin-gonic/gin"

var PORT string = ":8080"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"asd":     "dsa",
		})
	})
	r.Run(PORT) // listen and serve on 0.0.0.0:8080
}
