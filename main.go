package main

import (
	"52weeks/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/challenges", getChallenges)
	r.Run()
}
func getChallenges(c *gin.Context) {
	var challenges = []models.Challenge{
		{ID: "1", Name: "Férias de verão", Description: "Férias de verão com destino ao nordeste", TargetValue: 10000, Progress: 50},
		{ID: "2", Name: "Playstation 5", Description: "Console novo 2025", TargetValue: 5000, Progress: 20},
		{ID: "3", Name: "Playstation 5", Description: "Console novo 2025", TargetValue: 5000, Progress: 45},
	}
	c.JSON(200, challenges)
}
