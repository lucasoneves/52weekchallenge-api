package main

import (
	"52weeks/internal/data"
	"52weeks/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadChallenges()

	router := gin.Default()
	router.GET("/challenges", handler.GetChallenges)
	router.GET("/challenges/:id", handler.GetChallengeByID)
	router.PUT("/challenges/:id", handler.UpdateChallenge)
	router.DELETE("/challenges/:id", handler.DeleteChallenge)
	router.POST("/challenges", handler.CreateChallenge)
	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Olá Mundo"})
	})

	router.Run()
}
