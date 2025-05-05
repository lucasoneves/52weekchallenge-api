package main

import (
	"52weeks/models"

	"github.com/gin-gonic/gin"
)

var challenges = []models.Challenge{
	{ID: "1", Name: "Férias de verão", Description: "Férias de verão com destino ao nordeste", TargetValue: 10000, Progress: 50},
	{ID: "2", Name: "Playstation 5", Description: "Console novo 2025", TargetValue: 5000, Progress: 20},
	{ID: "3", Name: "Playstation 5", Description: "Console novo 2025", TargetValue: 5000, Progress: 45},
}

func main() {
	router := gin.Default()
	router.GET("/challenges", getChallenges)
	router.POST("/challenges", createChallenge)
	router.Run()
}
func getChallenges(c *gin.Context) {

	c.JSON(200, challenges)
}

func createChallenge(c *gin.Context) {
	var newChallenge models.Challenge
	if err := c.ShouldBindJSON(&newChallenge); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	challenges = append(challenges, newChallenge)
	c.JSON(200, newChallenge)
}
