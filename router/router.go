package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router using gin default configs
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// start API using 8080 door for default
	router.Run(":8080")
}
