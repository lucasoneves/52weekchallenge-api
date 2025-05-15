package main

import (
	"52weeks/internal/data"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadChallenges()

	router := gin.Default()
	router.GET("/challenges", getChallenges)
	router.GET("/challenges/:id", getChallengeByID)
	router.PUT("/challenges/:id", updateChallenge)
	router.DELETE("/challenges/:id", deleteChallenge)
	router.POST("/challenges", createChallenge)
	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Olá Mundo"})
	})

	router.Run()
}
