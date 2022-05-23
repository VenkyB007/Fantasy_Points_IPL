package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/sodhi_sowji", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"sodhi fellow": "sowjanya",
		})
	})

	r.Run()
}
